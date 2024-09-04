package solutions

func SumConcurrentUsingChannels(input []int) int {
	inputCh := make(chan int)
	resultCh := make(chan int)

	go func() {
		var sum int
		for {
			// `ok` indicates if the channel is still open or if it has been closed.
			num, ok := <-inputCh
			if !ok {
				break
			}
			sum += num
		}
		resultCh <- sum
	}()

	for _, number := range input {
		inputCh <- number
	}

	// Closing of the channel will result in a `break` in the `for` loop.
	close(inputCh)

	// Wait until the result comes out of the result channel, then return the result.
	return <-resultCh
}
