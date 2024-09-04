package worker

import (
	solution "github.com/minitauros/go-concurrency-training/course/5_worker/solutions"
)

func handleWork(concurrency int, workCh chan func()) {
	// Make sure that work coming in on the work channel is handled with the given concurrency.
	// Assume that the work channel will be closed as soon as all work has been handled.
	solution.PermanentRoutines(concurrency, workCh)
}
