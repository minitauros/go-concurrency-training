package solutions

import (
	"sync"
)

func SumConcurrentUsingWaitGroup(input []int) int {
	wg := &sync.WaitGroup{}

	var sum int
	for _, num := range input {
		wg.Add(1)
		go func() {
			sum += num
			wg.Done()
		}()
	}

	wg.Wait()
	return sum
}
