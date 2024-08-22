package rate_limiter

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	var numCalls atomic.Int64
	start := time.Now()
	limit(func() bool {
		numCalls.Add(1)
		if numCalls.Load() == 100 {
			return false
		}
		return true
	}, time.Millisecond)
	elapsed := time.Now().Sub(start)

	// Since we stop the limiter after 100 calls, and the rate limit is 1ms, the total time elapsed must be >=100ms.
	if elapsed < 100*time.Millisecond {
		t.Errorf("100 executions were reached in %v, which is less than the expected >=100ms", elapsed)
	}

	// It is expected that there is some overhead in sending things over channels. We want a callback call
	// no more than once per given rate, but not exactly once per given rate.
	// 200ms is an arbitrary number. This test is flaky. A slow machine might not be able to execute as many
	// callbacks as a fast machine.
	if elapsed > 200*time.Millisecond {
		t.Error("100 executions took longer than expected, even when accounting for overhead")
	}
}
