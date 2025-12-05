# Exercise

Work will be sent on the provided work channel. Make sure that this work is executed with the given concurrency
(e.g. if concurrency is 3, we want 3 pieces of work to be handled concurrently). You can stop listening for work as 
soon as the channel is closed.

You can test your solution by running the test:

```
go test -v .
```

## Detecting the closing of a channel

You can detect if a channel is closed in multiple ways.

```go
val, ok := <-myChan
```

`ok` will be `false` if the channel was closed.

```go
for val := range myChan {
	// Do something.
}
```

The `for` loop will automatically stop once `myChan` is closed.

## AI

* Explain how the worker pool pattern works in Go, the programming language.