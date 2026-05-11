package solutions

import (
	"sync"
)

func SumConcurrentUsingMutex(input []int) int {
	mux := &sync.Mutex{}
	wg := sync.WaitGroup{}

	var sum int
	for _, num := range input {
		wg.Go(func() {
			mux.Lock()
			sum += num
			mux.Unlock()
		})
	}

	wg.Wait()
	return sum
}
