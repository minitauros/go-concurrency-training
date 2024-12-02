package kafka

import (
	"errors"
	"maps"
	"slices"
	"sync"
	"sync/atomic"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

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
	wrappedConsumer *kafka.Consumer
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
	consumer *kafka.Consumer,
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
			event := c.wrappedConsumer.Poll(1000)
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

	// Stop consuming messages.
	close(c.stopCh)

	// Stop all routines that are handling work.
	for _, ch := range c.workChannels {
		close(ch)
	}
	for _, h := range c.handlers {
		h.Stop()
	}

	// Wait for all the work to finish.
	c.gracefulShutdownWg.Wait()

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
	c.workChannels[partition] = make(chan *kafka.Message)
	c.gracefulShutdownWg.Add(1)

	go func() {
		defer c.gracefulShutdownWg.Done()

		for msg := range c.workChannels[partition] {
			if err := c.handlers[*msg.TopicPartition.Topic].Handle(msg.Value); err != nil {
				c.errHandler(err, c.Stop)
			}
			<-c.concurrencyCh
		}
	}()
}
