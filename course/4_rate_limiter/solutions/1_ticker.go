package solutions

import (
	"time"
)

func LimitUsingTicker(callback func() bool, rate time.Duration) {
	for {
		// There is overhead in sending things over a channel, which may add a fraction of a millisecond of
		// overhead on each send.
		<-time.Tick(rate)
		if !callback() {
			// Clean up goroutine.
			return
		}
	}
}
