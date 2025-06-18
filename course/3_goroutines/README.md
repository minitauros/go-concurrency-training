## Task

Modify `sumConcurrent` in such a way that it concurrently sums numbers. To test your result, you can run the test:

```
go test -v .
```

Note that the test will also pass if you do not use goroutines, but for learning's sake, I recommend to try having the
test pass while using goroutines.

Goroutines are easy to start. Just type `go` in front of a function call, for example:

```go
// Do something in a goroutine.
go doSomething()

// Achieves the same.
go func() {
	// Do something.
}()
```
