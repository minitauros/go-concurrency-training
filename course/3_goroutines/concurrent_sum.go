package concurrency

import (
	"github.com/minitauros/go-concurrency-training/course/3_goroutines/solutions"
)

func sumConcurrent(input []int) int {
	return solutions.SumConcurrentUsingAtomic(input)
}
