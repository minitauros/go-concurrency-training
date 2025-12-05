# Exercise

In this exercise you will find the same code as in the warmup exercise. However, the test in the warmup exercise is
broken. It only covers the case where the input is `3`, and it doesn't even cover it correctly. Good tests cover all
possible routes that a code path can take.

It is possible to create a separate test function for each of those cases, but for this exercise we'll use `t.Run()` to
distinguish between test cases. For example:

```go
t.Run("If input is lower than 2", func (t) {
    // Test case implementation.
})

// ...other test cases
```

The task is to implement all possible test cases using `t.Run()` to distinguish between cases. To see if your tests
work, tests can be run as follows.

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