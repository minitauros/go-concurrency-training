# Task

In this exercise you will find the same code as in the warmup exercise. However, the test in the warmup exercise is broken. It doesn't cover all test cases.

It is possible to create a function for each test case, but in this case, we'll use `t.Run()` to distinguish between test cases.

Tests can be run as follows.

```
go test -v .
```

Where:

* **-v** = Verbose
* **.** = The directory in which to run the tests. In this case the current directory.


Validate that you have implemented the solution correctly by running the following command.

```
task test-solution
```

This will run a mutation testing tool. Mutation testing is a strategy in which the source code is being "maliciously" modified on purpose to see if that changes anything in the behavior of the tests. If the tests are written correctly and all cases are covered, the tests should start failing if the source code is modified. If the tests don't fail, you have very likely forgotten to test one or multiple cases.

## AI

* Explain what mutation testing is in software engineering.