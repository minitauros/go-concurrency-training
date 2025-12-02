package worker

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestGracefulShutdown(t *testing.T) {
	var workStartedCount atomic.Int64
	var workDoneCount atomic.Int64
	var workStartedBeforeShutdownCount int64
	var workDoneAfterShutdownCount int64

	w := newWorker()
	w.workStarted = func() {
		workStartedCount.Add(1)
	}

	stopCh := make(chan struct{})
	workCh := make(chan work)
	go w.performWork(4, workCh)
	go func() {
		for {
			select {
			case <-stopCh:
				return
			default:
				workCh <- func() {
					// Sleep to simulate work being done.
					time.Sleep(1 * time.Millisecond)
					workDoneCount.Add(1)
				}
			}
		}
	}()

	// Allow some work to be processed.
	time.Sleep(20 * time.Millisecond)

	// First stop sending work to prevent sending work on a closed channel.
	close(stopCh)

	workStartedBeforeShutdownCount = workStartedCount.Load()
	w.shutDown()
	workDoneAfterShutdownCount = workDoneCount.Load()

	fmt.Printf("workStartedBeforeShutdownCount: %+#v\n", workStartedBeforeShutdownCount)
	fmt.Printf("workFinishedAfterShutdownCount: %+#v\n", workDoneAfterShutdownCount)

	// Because of the implementation it is not safe to assert that work started EQUALS work done.
	// It is possible that the worker is reading work from the work channel, then the stop channel is closed (and the
	// "number of work started" is read), and then the work that has been taken from the work channel before stop was
	// called is still going to be executed, making the total amount of work executed higher than the amount of work we
	// just read has been started.
	if workDoneAfterShutdownCount < workStartedBeforeShutdownCount {
		t.Errorf("work started: %d, work finished: %d (expected to have as least as much work finished as started)", workStartedBeforeShutdownCount, workDoneAfterShutdownCount)
	}
}
