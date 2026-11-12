# Exercise

One of the most popular testing frameworks is [testify](https://github.com/stretchr/testify). It can be installed using
following command.

```shell
go get github.com/stretchr/testify
```

Since Testify is already being used in this project, installation probably did nothing, but that's fine, since we're
practicing Go here anyway, and running `go get` is some sort of practice.

In this exercise we're going to use Testify's `assert` package to assert that our code does what it is supposed to do.
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

Testify also has a `require` package with the same functions. The difference is what happens when an assertion fails:
`assert` marks the test as failed and continues, while `require` stops the test immediately. Use `require` when the
rest of the test makes no sense after a failure, for example when a setup step returns an error, or when a value that
the next lines dereference is `nil`:

```go
func TestFoo(t *testing.T) {
    user, err := repo.GetUser(ctx, 123)
    require.NoError(t, err)  // Stops here if err != nil.
    require.NotNil(t, user)  // Stops here if user == nil, so the next line can't panic.
    assert.Equal(t, "foo", user.Username)
}
```

`require` stops the test by calling `t.FailNow()`, which only works from the goroutine that runs the test. In other
goroutines, use `assert`.

First, implement the methods of `Calculator` in calculator.go in this directory. Then, using the `assert` package,
write tests for the implementations. The `SpecialSub()` methods has already been implemented and tested as an example.

To see if your tests work, tests can be run as follows.

```shell
go test -v .
```

When you are done, validate that you have implemented the solution correctly by running the following command.

```shell
task test-solution
```

## AI Prompts

* Show me an example of table driven tests in Go, the programming language.