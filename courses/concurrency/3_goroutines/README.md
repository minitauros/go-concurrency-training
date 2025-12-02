## Task

Modify `sumConcurrent` in such a way that it concurrently sums numbers. To test your result, you can run the test:

```
go test -v .
```

Note that the test will also pass if you do not use goroutines (i.e. if you don't touch the current implementation), 
but for learning's sake, I recommend to try achieving the same result using goroutines.

Goroutines are easy to start. Just type `go` in front of a function call, for example:

```go
// Do something in a goroutine.
go doSomething()

// Achieves the same.
go func() {
	// Do something.
}()
```

You will quickly notice that multiple goroutines writing to the same variable causes memory problems. You will also
likely run into the problem of the program returning before all goroutines have finished. To solve these issues, see 
the AI section.

## AI

* Explain how the atomic package works in Go, the programming language. Specifically atomic.Int64.
* Explain how a mutex works in Go, the programming language.
* Explain how a wait group works in Go, the programming language.
* Explain how I can use channels to concurrently sum up numbers in Go, the programming language.
