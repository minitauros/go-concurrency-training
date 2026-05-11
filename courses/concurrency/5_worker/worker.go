package worker

func handleWork(concurrency int, workCh chan func()) {
	// Make sure that work coming in on the work channel is handled with exactly the given concurrency.
	// For example, if the given concurrency is 3, execute no more than 3 pieces of work concurrently.
	// The work channel is automatically closed (by the test) as soon as all work has been handled.
}
