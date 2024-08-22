package worker

// Work is the work to be executed.
type work func()

type worker struct {
	// workStarted is used in the test to assert that the shutdown was graceful.
	// Please do not remove.
	workStarted func()
}

func newWorker() *worker {
	// The test will use this function to test your implementation. Please implement it - if needed - correctly.
	// You do not need to add the before and after shutdown functions; they will be set by the test.
	return &worker{}
}

// performWork will perform work coming on on the given channel, with the given concurrency.
func (w *worker) performWork(concurrency int, workCh chan work) {
	concurrencyCh := make(chan struct{}, concurrency)

	for {
		concurrencyCh <- struct{}{}
		work, ok := <-workCh
		if !ok {
			return
		}

		w.workStarted() // Call right before starting the goroutine in order for the test to work properly.
		go func() {
			work()
			<-concurrencyCh
		}()
	}
}

// shutDown will be called when we want to gracefully shut the worker down.
func (w *worker) shutDown() {

}
