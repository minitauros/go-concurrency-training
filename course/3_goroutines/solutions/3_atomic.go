package solutions

import (
	"sync"
	"sync/atomic"
)

func SumConcurrentUsingAtomic(input []int) int {
	var sum atomic.Int64
	var wg sync.WaitGroup

	for _, num := range input {
		wg.Add(1)
		go func() {
			sum.Add(int64(num))
			wg.Done()
		}()
	}

	wg.Wait()
	return int(sum.Load())
}
