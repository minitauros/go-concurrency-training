package main

import (
	"fmt"
	"time"
)

func main() {
	vals := []string{"mary", "had", "a", "little", "lamb", "heeya", "heeya", "ho"}

	for _, val := range vals {
		fmt.Print(val, " ")
	}

	fmt.Print("\n\n")

	for _, val := range vals {
		go fmt.Print(val, " ")
	}

	// Give goroutines time to finish.
	time.Sleep(10 * time.Millisecond)
}
