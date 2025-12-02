package solutions

// SumConcurrentUsingMadness is an implementation that is not Sparta; it is madness!
// To a first time Go writer it may seem like this would work, but it doesn't.
func SumConcurrentUsingMadness(input []int) int {
	var sum int
	for _, number := range input {
		go func() {
			sum += number
		}()
	}
	return sum
}
