package concurrency

func sumConcurrent(input []int) int {
	var total int
	for _, val := range input {
		total += val
	}
	return total
}
