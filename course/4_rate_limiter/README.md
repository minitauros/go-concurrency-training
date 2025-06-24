## Task

Write the implementation of `limit()` in limit.go in such a way that `callback` is called not more than once per
`rate` (e.g. not more than once per millisecond). As soon as the call to `callback` returns `false`, the function must
return/stop.

Validate that you have implemented the solution correctly by running the test:

```
go test -v .
```

You can use `time.Ticker` to ensure that something happens at a given interval. View
examples [here](https://gobyexample.com/tickers).

## AI

* Explain how time.Ticker works in Go, the programming language.