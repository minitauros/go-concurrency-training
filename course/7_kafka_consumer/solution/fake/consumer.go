package fake

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/minitauros/go-concurrency-training/course/7_kafka_consumer/pkg/kafka"
)

type Consumer struct {
	client      *kafka.Client
	concurrency int
	numRead     atomic.Int64
	maxNumRead  int
	// sleepTime indicates how long to sleep after a message is handled.
	sleepTime time.Duration
	stopCh    chan struct{}
	stopWg    sync.WaitGroup
}

func (c *Consumer) Start() error {
	for i := 0; i < c.concurrency; i++ {
		c.stopWg.Add(1)
		go func() {
			defer c.stopWg.Done()

			select {
			case <-c.stopCh:
				return
			default:
				msg := c.client.ReadMessage()
				fmt.Println("read message: ", msg)

				newVal := c.numRead.Add(1)
				if int(newVal) == c.maxNumRead {
					return
				}

				time.Sleep(c.sleepTime)
			}
		}()
	}
	return nil
}

func (c *Consumer) Stop() error {
	close(c.stopCh)
	c.stopWg.Wait()
	return nil
}
