## Task

Modify `sumConcurrent` in such a way that it concurrently sums numbers. To test your result, you can run the test, which
should also pass when you implement the solution correctly. You can run the test as follows:

```
go test -v .
```

Where `-v` means "verbose" and `.` is the current directory.

Goroutines are easy to start. Just type `go` in front of a function call, for example:

```go
// Do something in a goroutine.
go doSomething()

// Achieves the same.
go func() {
	// Do something.
}()
```
