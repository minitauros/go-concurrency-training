package main

import (
	"fmt"
	"time"
)

func main() {
	vals := []string{"mary", "had", "a", "little", "lamb", "heeya", "heeya", "ho"}

	fmt.Println("printAll():")
	printAll(vals)

	fmt.Print("\n\n")
	fmt.Println("printAllConcurrent():")
	printAllConcurrent(vals)

	// Give goroutines time to finish.
	time.Sleep(100 * time.Millisecond)
}

func printAll(vals []string) {
	for _, val := range vals {
		fmt.Print(val, " ")
	}
}

func printAllConcurrent(vals []string) {
	for _, val := range vals {
		go fmt.Print(val, " ")
	}
}
