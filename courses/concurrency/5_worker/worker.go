package worker

func handleWork(concurrency int, workCh chan func()) {
	// Make sure that work coming in on the work channel is handled with the given concurrency.
	// The work channel is automatically closed (by the test) as soon as all work has been handled.
}
