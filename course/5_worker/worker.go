package worker

func handleWork(concurrency int, workCh chan func()) {
	// Make sure that work coming in on the work channel is handled with the given concurrency.
	// Assume that the work channel will be closed as soon as all work has been handled.
}
