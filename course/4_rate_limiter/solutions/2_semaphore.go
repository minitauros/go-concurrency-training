package solutions

import (
	"time"
)

func LimitUsingSemaphore(callback func() bool, rate time.Duration) {
	limiter := make(chan struct{}, 10) // Semaphore/limiter with burst of 10.
	stopCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-stopCh:
				return
			case <-time.Tick(rate):
				<-limiter
			}
		}
	}()

	for {
		// Note that this does not guarantee that EXACTLY every millisecond a tick happens.
		// There is overhead in sending things over a channel, which may add a fraction of a millisecond of
		// overhead on each send.
		limiter <- struct{}{}
		if !callback() {
			// Clean up goroutine.
			close(stopCh)
			return
		}
	}
}
