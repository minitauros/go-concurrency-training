package fake

import (
	"testing"
	"time"

	"github.com/minitauros/go-concurrency-training/course/7_kafka_consumer/pkg/kafka"
)

func Test_Consumer_Start(t *testing.T) {
	consumer := &Consumer{
		client:      &kafka.Client{},
		concurrency: 2,
		maxNumRead:  2,
		sleepTime:   50 * time.Millisecond,
		stopCh:      make(chan struct{}),
	}

	_ = consumer.Start()

	// Allow goroutines to start.
	time.Sleep(time.Millisecond)

	// Since we are sleeping 50 milliseconds after each message, have a concurrency of 2, and are reading max 2 messages,
	// we expect 2 messages to be read immediately.
	// If concurrency is not working, we would expect the consumption to take at least 100 milliseconds.
	if consumer.numRead.Load() < 2 {
		t.Error("failed to read messages concurrently")
	}

	// Clean up goroutines.
	_ = consumer.Stop()
}

func Test_Consumer_Stop(t *testing.T) {
	consumer := &Consumer{
		client:      &kafka.Client{},
		concurrency: 1,
		sleepTime:   50 * time.Millisecond,
		stopCh:      make(chan struct{}),
	}

	now := time.Now()
	_ = consumer.Start()

	// Allow goroutines to start.
	time.Sleep(time.Millisecond)

	_ = consumer.Stop()

	// We are sending one message, and we sleep 50 millisecond after each message.
	// That means Stop() may only return after those 50 milliseconds.
	if time.Since(now) < consumer.sleepTime {
		t.Error("failed to wait until all messages were processed")
	}
}
