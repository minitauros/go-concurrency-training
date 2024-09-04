package solutions

import (
	"sync"
)

func SumConcurrentUsingMutex(input []int) int {
	mux := &sync.Mutex{}
	wg := sync.WaitGroup{}

	var sum int
	for _, num := range input {
		wg.Add(1)
		go func() {
			mux.Lock()
			sum += num
			mux.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	return sum
}
