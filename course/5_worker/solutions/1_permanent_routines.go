package solution

func PermanentRoutines(concurrency int, workCh chan func()) {
	for i := 0; i < concurrency; i++ {
		go func() {
			// It is possible to use `for .. range` to read from channels. The loop will automatically break as soon
			// as the channel is closed.
			for work := range workCh {
				work()
			}
		}()
	}
}
