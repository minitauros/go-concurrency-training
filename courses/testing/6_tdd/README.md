# Exercise

In this exercise we'll have a minimal introduction to test driven development, which is when you don't start with
writing the implementation, but with writing the tests. You write a test, run it, expect it to fail, then you implement
what the test tests, rerun the test, and if it passes, you keep this process going until all tests have been written,
all tests pass, and the implementation is complete.

You may notice that the tests in this package do not live in the `tdd` package, but in the `tdd_test` package. Go allows
you to put tests in a package that ends with `_test`, and even if the tests are in the same directory, they are treated
as if they are in a separate package.

In this case, that is necessary to prevent an import cycle where the tests in package `tdd` import the mocks from
`mocks`, while the mocks in `mocks` import `tdd` because the mocks need to return `tdd.User`.

Now, as to what to do, look at the tests in this package and write the implementation. 

Your work is complete as soon as all the tests pass. You may find it handy to use the `failfast` flag with the tests so that the tests stop running as soon as a test is found that fails. This way your console won't be full with output of all failing tests.

```shell
go test -v -failfast .
```

## AI Prompts

* Explain to me what test driven development (TDD) is and give me an example in Go.