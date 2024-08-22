package solutions

// channels1 returns the message that is sent on messageCh.
func channels1(messageCh chan string) string {
	return <-messageCh
}

// channels2 ensures that whatever is received on the input channel is sent into the output channel.
func channels2(inputCh chan string, outputCh chan string) {
	for val := range inputCh {
		outputCh <- val
	}
}

// channels3 returns a channel on which all the given input will be sent.
func channels3(input []string) chan string {
	ch := make(chan string)

	go func() {
		for _, in := range input {
			ch <- in
		}
	}()

	return ch
}
