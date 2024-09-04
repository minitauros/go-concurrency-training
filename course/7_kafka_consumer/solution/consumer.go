package solution

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	consumer "github.com/minitauros/go-concurrency-training/course/7_kafka_consumer"
	"github.com/segmentio/kafka-go"
)

// Ensure interface compliance.
var _ consumer.Consumer = (*Consumer)(nil)

type HandlerFn func(msg []byte)

type Consumer struct {
	reader   *kafka.Reader
	handlers map[string]HandlerFn
	// concurrency defines how many routines will be working to handle the work.
	concurrency int
	// gracefulShutdownWg is used for graceful work shutdown.
	gracefulShutdownWg *sync.WaitGroup
	// workCh is the channel on which we will send messages to be handled.
	workCh chan kafka.Message
	// workerReadyCh is used to make sure that we don't pick up more work than we can handle.
	workerReadyCh chan struct{}
	// criticalErrCh is the channel that errors are sent on that should stop the consumer.
	criticalErrCh chan error
	// stopCh is a channel that will be closed once the consumer is stopped, so that we can orchestrate
	// a full shutdown of everything.
	stopCh         chan struct{}
	alreadyStopped atomic.Bool
}

func NewConsumer(reader *kafka.Reader, concurrency int) *Consumer {
	return &Consumer{
		reader:             reader,
		handlers:           make(map[string]HandlerFn),
		concurrency:        concurrency,
		workerReadyCh:      make(chan struct{}, concurrency),
		gracefulShutdownWg: &sync.WaitGroup{},
		workCh:             make(chan kafka.Message),
		criticalErrCh:      make(chan error, concurrency),
		stopCh:             make(chan struct{}),
	}
}

func (c *Consumer) SetHandler(topicName string, handler HandlerFn) {
	c.handlers[topicName] = handler
}

func (c *Consumer) Start() error {
	c.startWorkers()

	for {
		select {
		case <-c.stopCh:
			return nil
		case err := <-c.criticalErrCh:
			stopErr := c.Stop()
			if stopErr != nil {
				return errors.Join(err, stopErr)
			}
			return err
		case c.workerReadyCh <- struct{}{}: // Only try to fetch a new message from Kafka if there is a worker ready to handle it.
			msg, err := c.reader.FetchMessage(context.Background())
			if err != nil {
				c.criticalErrCh <- err
				continue
			}
			c.workCh <- msg
		}
	}
}

func (c *Consumer) startWorkers() {
	for i := 0; i < c.concurrency; i++ {
		// Add one to the wait group, so that when we shut down we can wait for each goroutine to finish.
		c.gracefulShutdownWg.Add(1)

		go func() {
			// Make sure that when the routine finishes we signal to the wait group that it's done.
			defer c.gracefulShutdownWg.Done()

			for {
				select {
				case <-c.stopCh:
					return
				case msg := <-c.workCh:
					err := c.handleMessage(msg)
					if err != nil {
						c.criticalErrCh <- err
					}

					// Take a token from the pool so that the next message can be retrieved from Kafka.
					<-c.workerReadyCh
				}
			}
		}()
	}
}

func (c *Consumer) handleMessage(msg kafka.Message) error {
	handler, ok := c.handlers[msg.Topic]
	if !ok {
		return ErrNoHandlerForTopic{TopicName: msg.Topic}
	}
	handler(msg.Value)
	return c.reader.CommitMessages(context.Background(), msg)
}

func (c *Consumer) Stop() error {
	// Set the value to true. If the value before the swap was also `true`, don't stop again; just return `nil`.
	if c.alreadyStopped.Swap(true) {
		return nil
	}

	// Closing a channel will send a signal to all listeners.
	// In the following example `val` will be an empty value and `ok` will be false as soon as the channel gets closed.
	//
	// val, ok := <-c.stopCh
	close(c.stopCh)

	// Wait for all the goroutines to finish their work.
	c.gracefulShutdownWg.Wait()

	return c.reader.Close()
}

// ErrNoHandlerForTopic is used when there is no handler defined for a topic that a message was consumed from.
type ErrNoHandlerForTopic struct {
	TopicName string
}

// Error satisfies the error interface.
func (e ErrNoHandlerForTopic) Error() string {
	return fmt.Sprintf("no handler for topic %s", e.TopicName)
}
