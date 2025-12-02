package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Duration(i) * time.Millisecond)
			if i == 2 {
				close(done)
			}
		}
	}()

	go func() {
		<-time.After(time.Millisecond)
		fmt.Println("wow")
	}()

	<-done
	fmt.Println("done")
}
