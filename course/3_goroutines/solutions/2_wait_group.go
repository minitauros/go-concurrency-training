package solutions

import (
	"sync"
)

// SumConcurrentUsingWaitGroup solves the problem of the program returning before all input was processed, but does
// not protect against concurrent memory access.
func SumConcurrentUsingWaitGroup(input []int) int {
	wg := sync.WaitGroup{}

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
