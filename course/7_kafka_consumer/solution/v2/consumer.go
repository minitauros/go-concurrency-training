package kafka

import (
	"errors"
	"maps"
	"slices"
	"sync"
	"sync/atomic"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const PollInterval = 1000 * time.Millisecond

type WrappedConsumer interface {
	SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) (err error)
	Poll(timeoutMs int) (event kafka.Event)
	Close() (err error)
}

// Handler handles Kafka messages.
type Handler interface {
	// Handle handles the given message. If an error is returned, it is passed to the error handler and the consumer is
	// stopped. Note that this function may be called concurrently and ought thus to be concurrency safe.
	Handle(msg []byte) error
	// Stop stops the handler. This is mainly useful for concurrent handlers. When the consumer is stopped, it will also
	// wait for all handlers to stop.
	Stop()
}

type Consumer struct {
	wrappedConsumer WrappedConsumer
	handlers        map[string]Handler
	stopCh          chan struct{}
	stoppedCh       chan struct{}
	// errHandler handles errors. Note that the error handler may be called concurrently. Calling `stop` stops the
	// consumer and is equivalent to calling `Stop()` on the consumer.
	errHandler func(err error, stop func() error)
	// workChannels is a map that contains a channel to send work on, indexed by partition. That means that the work
	// for topic A and topic B's partition one will both end up in the same work channel. That is not a problem,
	// because this still means that there can be no more than one message of a partition of any given topic in memory
	// at any given time.
	workChannels map[int]chan *kafka.Message
	// concurrencyCh is used for controlling the concurrency while handling messages.
	concurrencyCh chan struct{}
	// gracefulShutdownWg is used for graceful work shutdown.
	gracefulShutdownWg sync.WaitGroup
	alreadyStopped     atomic.Bool
}

func NewConsumer(
	consumer WrappedConsumer,
	handlers map[string]Handler,
	errHandler func(err error, stop func() error),
	concurrency int,
) (*Consumer, error) {
	topicNames := slices.Collect(maps.Keys(handlers))
	if len(topicNames) == 0 {
		return nil, errors.New("no topics to consume given")
	}
	if slices.Contains(topicNames, "") {
		return nil, errors.New("topic name cannot be empty string")
	}

	err := consumer.SubscribeTopics(topicNames, nil)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		stopCh:          make(chan struct{}),
		stoppedCh:       make(chan struct{}),
		wrappedConsumer: consumer,
		handlers:        handlers,
		errHandler:      errHandler,
		workChannels:    make(map[int]chan *kafka.Message),
		concurrencyCh:   make(chan struct{}, concurrency),
	}, nil
}

func (c *Consumer) Start() {
	for {
		select {
		case <-c.stopCh:
			return
		case c.concurrencyCh <- struct{}{}:

			event := c.wrappedConsumer.Poll(int(PollInterval.Milliseconds()))
			if event == nil {
				<-c.concurrencyCh
				continue
			}

			switch e := event.(type) {
			case *kafka.Message:
				if _, ok := c.workChannels[int(e.TopicPartition.Partition)]; !ok {
					c.startWorkHandler(int(e.TopicPartition.Partition))
				}
				c.workChannels[int(e.TopicPartition.Partition)] <- e

				// No need to take one from the concurrency channel here; the goroutine that is handling the work for
				// this message's partition will do that.

				// TODO Commit async.
			case kafka.Error:
				c.errHandler(e, c.Stop)
				<-c.concurrencyCh
			default:
				// ignore the rest
				<-c.concurrencyCh
			}
		}
	}
}

func (c *Consumer) Stop() error {
	if c.alreadyStopped.Swap(true) {
		// The value before the swap was already `true`, which means `Stop()` has already been called at least once.
		return nil
	}

	close(c.stopCh)

	// Wait for all the workers to finish.
	c.gracefulShutdownWg.Wait()

	// Stop each handler individually for the work may be handled concurrently and this allows us to wait for that to
	// finish.
	for _, h := range c.handlers {
		h.Stop()
	}

	if err := c.wrappedConsumer.Close(); err != nil {
		return err
	}

	// Signal to the outside world that the consumer has stopped.
	close(c.stoppedCh)

	return nil
}

func (c *Consumer) Stopped() <-chan struct{} {
	return c.stoppedCh
}

func (c *Consumer) startWorkHandler(partition int) {
	workCh := make(chan *kafka.Message)
	c.workChannels[partition] = workCh
	c.gracefulShutdownWg.Add(1)

	go func() {
		defer c.gracefulShutdownWg.Done()

		for {
			select {
			case <-c.stopCh:
				return
			case msg := <-workCh:
				if err := c.handlers[*msg.TopicPartition.Topic].Handle(msg.Value); err != nil {
					// The error handler must be called in a goroutine, just in case it wants to call Stop.
					// Stop will wait for gracefulShutdownWg, and if the call to the error handler is blocking,
					// the current goroutine will never return, and will thus never be able to signal gracefulShutdownWg
					// that it's done.
					//
					// Example:
					// Handler returns an error.
					// Error is passed to the consumer's error handler.
					// Error handler calls stop(), which is just a call to Stop() on the consumer.
					// Consumer waits for all the work to finish.
					// Work can not finish, because there is one worker closing the consumer, while closing is blocked
					// by waiting for all the work to finish, which will never happen because of the worker that is
					// trying to close the consumer.
					// Deadlock.
					go c.errHandler(err, c.Stop)
				}
				<-c.concurrencyCh
			}
		}
	}()
}
