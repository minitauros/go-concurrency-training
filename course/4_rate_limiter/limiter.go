package rate_limiter

import (
	"time"
)

func limit(callback func() bool, rate time.Duration) {
	for {
		if !callback() {
			return
		}
	}
}
