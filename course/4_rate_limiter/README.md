## Task

Write the function implementation so that `callback` is called not more than once per `rate` (e.g. not more than once
per millisecond), until `callback` returns `false`, after which the function completely stops.

Validate that you have implemented the solution correctly by running the test:

```
go test -v .
```

You can use `time.Ticker` to ensure that something happens at a given interval. View
examples [here](https://gobyexample.com/tickers).


## AI

* Explain how time.Ticker works in Go, the programming language.