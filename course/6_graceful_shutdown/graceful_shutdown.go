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

// performWork will perform work coming on the given channel, with the given concurrency.
func (w *worker) performWork(concurrency int, workCh chan work) {
	// Token pool to make sure that no more than `concurrency` goroutines can run at the same time.
	concurrencyCh := make(chan struct{}, concurrency)

	for {
		// Put a token into the pool to reserve a concurrency slot for doing work.
		concurrencyCh <- struct{}{}

		work, ok := <-workCh
		if !ok {
			// No work left to do.
			return
		}

		// Needed for the tests to pass. Please make sure that this is called before starting the goroutine in which
		// the work is performed in order for the test to work properly.
		w.workStarted()
		go func() {
			work()

			// Return token to pool so that a next piece of work can be handled.
			<-concurrencyCh
		}()
	}
}

// shutDown will be called when we want to gracefully shut the worker down.
// That means it will wait for all the running work to finish before returning.
func (w *worker) shutDown() {
	// Allow no more work and wait for all work to finish.
}
