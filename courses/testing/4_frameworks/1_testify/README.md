# Exercise

One of the most popular testing frameworks is [testify](https://github.com/stretchr/testify). It can be installed using
following command.

```shell
go get github.com/stretchr/testify
```

Since testify is already being used in this project, installation probably did nothing, but that's fine, since we're
practicing Go here anyway, and running `go get` is some sort of practice.

In this exercise we're going to use testify's `assert` package to assert that our code does what it is supposed to do.
First, import the `assert` package:

```go
import "github.com/stretchr/testify/assert"
```

`assert` can be used in two ways. You can either create an `assert` variable that contains all the assertion functions:

```go
func TestFoo(t *testing.T) {
    assert := assert.New(t)
    assert.Equal(1, 1)
}
```

Or you can call `assert` package functions directly, passing it `*testing.T`:

```go
func TestFoo(t *testing.T) {
    assert.Equal(t, 1, 1)
}
```

First, implement the methods of `Calculator` in calculator.go. Then, using the `assert` package, write tests for the 
implementations. The `SpecialSub()` methods has already been implemented and tested as an example.

To see if your tests work, tests can be run as follows.

```
go test -v .
```

When you are done, validate that you have implemented the solution correctly by running the following command.

```
task test-solution
```

## AI

* Show me an example of table driven tests in Go, the programming language...