package solution

func RoutinePerJob(concurrency int, workCh chan func()) {
	concurrencyCh := make(chan struct{}, concurrency)

	for {
		concurrencyCh <- struct{}{}
		work, ok := <-workCh
		if !ok {
			return
		}

		go func() {
			work()
			<-concurrencyCh
		}()
	}
}
