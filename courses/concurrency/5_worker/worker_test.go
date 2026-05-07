package worker

import (
	"sync"
	"testing"
	"time"
)

func TestHandleWork(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var work1DoneAt time.Time
	var work2DoneAt time.Time

	work1 := func() {
		time.Sleep(5 * time.Millisecond)
		work1DoneAt = time.Now()
		wg.Done()
	}

	work2 := func() {
		work2DoneAt = time.Now()
		wg.Done()
	}

	workCh := make(chan func())
	go handleWork(2, workCh)

	workCh <- work1
	workCh <- work2
	close(workCh)

	doneCh := make(chan struct{})
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	select {
	case <-time.After(10 * time.Millisecond):
		t.Error("not all work was done within a reasonable time")
	case <-doneCh:
		break
	}

	// If the handler is not concurrent, we would expect work1 to be done before work2.
	// In a concurrent setup, work2 is done first, because it can be done while work1 is sleeping.
	if work1DoneAt.Before(work2DoneAt) {
		t.Errorf("work1 was done (%v) before work2 (%v), which should not be possible if work is being performed concurrently", work1DoneAt, work2DoneAt)
	}
}
