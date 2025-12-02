package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	workCh := make(chan int)
	errCh := make(chan error)

	wg.Add(1)
	go func() {
		defer wg.Done()

		var i int
		for {
			select {
			case err := <-errCh:
				fmt.Println("error:", err)
			default:
				if i == 5 {
					close(workCh)
					return
				}
				fmt.Println("sending work")
				workCh <- i
				i++
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for work := range workCh {
			fmt.Println("doing work: ", work)
			err := doWork(work)
			if err != nil {
				fmt.Println("sending err")
				errCh <- err
			}
		}
	}()

	wg.Wait()
	fmt.Println("done")
}

func doWork(i int) error {
	if i == 2 {
		return errors.New("cannot handle 2")
	}
	return nil
}
