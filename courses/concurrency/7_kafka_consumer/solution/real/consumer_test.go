package kafka

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/minitauros/go-concurrency-training/courses/concurrency/7_kafka_consumer/solution/real/mocks"
	"github.com/minitauros/go-concurrency-training/courses/concurrency/7_kafka_consumer/solution/real/pkg/test"
	assert2 "github.com/stretchr/testify/assert"
)

type mockHandler struct {
	handle func(msg []byte) error
	stop   func()
}

func (m *mockHandler) Handle(msg []byte) error {
	if m.handle != nil {
		return m.handle(msg)
	}
	return nil
}

func (m mockHandler) Stop() {
	if m.stop != nil {
		m.stop()
	}
}

func Test_Consumer_Start(t *testing.T) {
	assert := assert2.New(t)

	t.Run("when the stop channel is closed, returns", func(t *testing.T) {
		c := &Consumer{
			stopCh: make(chan struct{}),
		}
		close(c.stopCh)

		doneCh := make(chan struct{})
		go func() {
			c.Start()
			close(doneCh)
		}()

		for {
			select {
			case <-doneCh:
				// Test passes if code reaches this point.
				return
			case <-time.After(5 * time.Millisecond):
				t.Error("closing stop channel did not make Start() return")
				return
			}
		}
	})

	t.Run("if there is no space on the concurrency channel, does not poll for messages", func(t *testing.T) {
		wrappedConsumer := mocks.NewMockWrappedConsumer(t)
		c := &Consumer{
			stopCh:          make(chan struct{}),
			concurrencyCh:   make(chan struct{}),
			wrappedConsumer: wrappedConsumer,
		}

		// Start the consumer and give it a little bit of time to consume and start polling.
		// The test will succeed if in all the time given to the consumer, it does not poll for messages.
		// The mock we generated will output an error if any calls are made that were not expected.
		go c.Start()
		time.Sleep(5 * time.Millisecond)
		close(c.stopCh)
	})

	t.Run("if there is space on the concurrency channel", func(t *testing.T) {
		var wrappedConsumer *mocks.MockWrappedConsumer
		var handler *mockHandler
		var c *Consumer

		topicName := test.String()
		setup := func() {
			wrappedConsumer = mocks.NewMockWrappedConsumer(t)
			handler = &mockHandler{}
			c = &Consumer{
				stopCh:          make(chan struct{}),
				concurrencyCh:   make(chan struct{}, 2),
				wrappedConsumer: wrappedConsumer,
				workChannels:    make(map[int]chan *kafka.Message),
				handlers: map[string]Handler{
					topicName: handler,
				},
			}
		}

		t.Run("if the consumed event is a Kafka message, calls the correct handler concurrently", func(t *testing.T) {
			setup()

			// We will generate two messages, same topic, different partition, because partitions are being handled
			// concurrently.
			//
			// We will sleep for x time in the handler.
			//
			// If the messages are NOT handled concurrently, we expect the program to have slept for x*2
			// (the full sleep time for each produced message).
			//
			// If they are handled concurrently, we expect the program to have slept for less than x*2 (the handlers all
			// sleep concurrently, so if the sleep time is 5ms, the total sleep time should be around 5ms even if there
			// are multiple handlers sleeping for 5ms each).

			sleepTime := 5 * time.Millisecond
			var numCalls atomic.Int64
			start := time.Now()
			var elapsed time.Duration

			// Since we have a concurrency of 2, return 2 messages for a different partition, which we expect to be
			// handled concurrently.
			wrappedConsumer.EXPECT().
				Poll(int(PollInterval.Milliseconds())).
				Once().
				Return(&kafka.Message{
					TopicPartition: kafka.TopicPartition{
						Topic:     &topicName,
						Partition: 1,
					},
				})
			wrappedConsumer.EXPECT().
				Poll(int(PollInterval.Milliseconds())).
				Once().
				Return(&kafka.Message{
					TopicPartition: kafka.TopicPartition{
						Topic:     &topicName,
						Partition: 2,
					},
				})

			// The code might be able to poll for more messages before we manage to stop the consumer.
			wrappedConsumer.EXPECT().
				Poll(int(PollInterval.Milliseconds())).
				Return(nil)

			handler.handle = func(msg []byte) error {
				time.Sleep(sleepTime)

				numCalls.Add(1)
				if numCalls.Load() == 2 {
					elapsed = time.Now().Sub(start)

					// Stop the consumer.
					close(c.stopCh)
				}

				return nil
			}

			c.Start()

			assert.Less(elapsed, sleepTime*2)
		})

		t.Run("if the consumed event is a Kafka error, passes it to the error handler", func(t *testing.T) {
			setup()

			kafkaErr := kafka.NewError(kafka.ErrBadMsg, test.String(), false)

			wrappedConsumer.EXPECT().
				Poll(int(PollInterval.Milliseconds())).
				Once().
				Return(kafkaErr)

			// The code might be able to poll for more messages before we manage to stop the consumer.
			wrappedConsumer.EXPECT().
				Poll(int(PollInterval.Milliseconds())).
				Return(nil)

			var receivedErr kafka.Error
			c.errHandler = func(err error, stop func() error) {
				receivedErr = err.(kafka.Error)

				// Stop the consumer.
				close(c.stopCh)
			}

			c.Start()

			assert.Equal(kafkaErr, receivedErr)
		})

		t.Run("if the consumed event is none of the above, does nothing", func(t *testing.T) {
			setup()

			wrappedConsumer.EXPECT().
				Poll(int(PollInterval.Milliseconds())).
				Return(nil)

			var msgHandled atomic.Bool
			var errHandled atomic.Bool

			handler.handle = func(msg []byte) error {
				msgHandled.Store(true)
				return nil
			}
			c.errHandler = func(err error, stop func() error) {
				errHandled.Store(true)
			}

			go c.Start()

			// Allow the consumer to process some events.
			time.Sleep(5 * time.Millisecond)

			close(c.stopCh)

			assert.Equal(false, msgHandled.Load())
			assert.Equal(false, errHandled.Load())
		})
	})
}

func Test_Consumer_Stop(t *testing.T) {
	assert := assert2.New(t)

	t.Run("if stopped before, does not stop again", func(t *testing.T) {
		wrappedConsumer := mocks.NewMockWrappedConsumer(t)
		c := &Consumer{
			concurrencyCh:   make(chan struct{}),
			wrappedConsumer: wrappedConsumer,
		}
		c.alreadyStopped.Store(true)

		err := c.Stop()
		assert.Nil(err)

		// The presence of the mock will ensure that no calls to that mock are being made, meaning that the consumer
		// has not done its regular stopping work.
	})

	t.Run("if not stopped before, correctly stops", func(t *testing.T) {
		wrappedConsumer := mocks.NewMockWrappedConsumer(t)

		var handlerStopped atomic.Bool
		handler := &mockHandler{
			stop: func() {
				handlerStopped.Store(true)
			},
		}
		c := &Consumer{
			stopCh:          make(chan struct{}),
			stoppedCh:       make(chan struct{}),
			concurrencyCh:   make(chan struct{}),
			wrappedConsumer: wrappedConsumer,
			handlers: map[string]Handler{
				test.String(): handler,
			},
		}

		wrappedConsumer.EXPECT().
			Close().
			Once().
			Return(nil)

		var stopChClosed atomic.Bool
		go func() {
			_, ok := <-c.stopCh
			if !ok {
				stopChClosed.Store(true)
			}
		}()

		var stoppedChClosed atomic.Bool
		go func() {
			_, ok := <-c.stoppedCh
			if !ok {
				stoppedChClosed.Store(true)
			}
		}()

		// Pretend we are doing some work.
		c.gracefulShutdownWg.Add(1)

		var hasStopped atomic.Bool
		var hasStopErr atomic.Bool
		go func() {
			err := c.Stop()
			if err != nil {
				hasStopErr.Store(true)
			}
			hasStopped.Store(true)
		}()

		// After a fraction of time, we assume that the goroutine we just spun up has been able to call Stop().
		// However, Stop() must wait for the wait group to be done, and we haven't signaled yet that it is,
		// // so Stop() must still be waiting.
		time.Sleep(5 * time.Millisecond)
		assert.False(hasStopped.Load())

		// Pretend all work is finished.
		c.gracefulShutdownWg.Done()

		// Allow the goroutine to finish.
		time.Sleep(5 * time.Millisecond)

		assert.True(hasStopped.Load())
		assert.True(handlerStopped.Load())
		assert.True(stopChClosed.Load())
		assert.True(stoppedChClosed.Load())
	})
}

func Test_Consumer_Stopped(t *testing.T) {
	assert := assert2.New(t)

	c := &Consumer{
		stoppedCh: make(chan struct{}),
	}

	ch := c.Stopped()

	var stopped atomic.Bool
	go func() {
		_, ok := <-ch
		if !ok {
			stopped.Store(true)
		}
	}()

	assert.False(stopped.Load())

	close(c.stoppedCh)

	// Allow the goroutine to finish.
	time.Sleep(5 * time.Millisecond)

	assert.True(stopped.Load())
}
