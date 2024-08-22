package solutions

import (
	"time"
)

func LimitUsingSemaphore(callback func() bool, rate time.Duration) {
	limiter := make(chan struct{}, 10) // Semaphore/limiter with burst of 10.

	// Fill the limiter so that a burst can happen.
	for i := 0; i < 10; i++ {
		limiter <- struct{}{}
	}

	stopCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-stopCh:
				return
			case <-time.Tick(rate):
				limiter <- struct{}{}
			}
		}
	}()

	for {
		// Note that this does not guarantee that EXACTLY every millisecond a tick happens.
		// There is overhead in sending things over a channel, which may add a fraction of a millisecond of
		// overhead on each send.
		<-limiter
		if !callback() {
			// Clean up goroutine.
			close(stopCh)
			return
		}
	}
}
