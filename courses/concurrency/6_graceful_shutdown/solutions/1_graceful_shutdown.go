package worker

import (
	"sync"
)

// Work is the work to be executed.
type work func()

type worker struct {
	stopCh      chan struct{}
	stopWg      sync.WaitGroup
	workStarted func()
}

func newWorker() *worker {
	return &worker{
		stopCh: make(chan struct{}),
	}
}

// performWork will perform work coming on on the given channel, with the given concurrency.
func (w *worker) performWork(concurrency int, workCh chan work) {
	concurrencyCh := make(chan struct{}, concurrency)

	for {
		concurrencyCh <- struct{}{}

		select {
		case <-w.stopCh:
			return
		case work, ok := <-workCh:
			if !ok {
				return
			}

			w.workStarted()
			w.stopWg.Add(1)
			go func() {
				work()
				w.stopWg.Done()
				<-concurrencyCh
			}()
		}
	}
}

// shutDown will be called when we want to gracefully shut the worker down.
func (w *worker) shutDown() {
	close(w.stopCh)
	w.stopWg.Wait()
}
