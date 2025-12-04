# Task

In this exercise you will find the same code as in the warmup exercise. However, the test in the warmup exercise is
broken. It doesn't cover all test cases.

It is possible to create a function for each test case, but for this exercise we'll use `t.Run()` to distinguish between
test cases. For example.

```go
t.Run("Some test case description", func (t) {
    // Test case implementation.
})
```

To see if your tests work, tests can be run as follows.

```
go test -v .
```

Where:

* **-v** = Verbose
* **.** = The directory in which to run the tests. In this case the current directory.

When you are done, validate that you have implemented the solution correctly by running the following command.

```
task test-solution
```

This will run a mutation testing tool that will run your test a number of times under different circumstances. If the
CLI shows green checkmarks only, you've successfully completed this exercise.

Mutation testing is a strategy in which the source code is being "maliciously" modified on purpose to see if that
changes anything in the behavior of the tests. If the tests are written correctly and all cases are covered, the tests
should start failing if the source code is modified. If the tests don't fail, it is very likely that you have forgotten
to test one or multiple cases.

## AI

* Explain what mutation testing is in software engineering.