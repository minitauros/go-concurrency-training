package concurrency

func sumConcurrent(input []int) int {
	var sum int
	for _, number := range input {
		sum += number
	}
	return sum
}
