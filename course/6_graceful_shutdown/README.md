## Task

Make sure that as soon as `shutDown()` returns, all work has finished executing. You can modify all the functions in graceful_shutdown.go, but please do not rename them, as the test relies on them existing.

Reading from a channel is done as follows:

```go
val, ok := <-someCh
```

Where `val` becomes the value sent on the channel, and `ok` indicates if the channel is still open.

You can test your solution by running the test:

```
go test -v .
```