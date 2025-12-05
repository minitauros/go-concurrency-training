# Task

Using a table driven format for tests can make it easier to write, read, and maintain tests. Where in a typical scenario
you would write out each test case as an individual test, when using table driven tests, you write an array of test
inputs and expected outputs and run them all in a loop. For example:

```go
type input struct {
	someArg string
}
type expected struct {
	someReturnArg error
}
testCases := []struct {
	description string
	input       input
	expected    expected
}{
    {
        description: "If input is x, returns y",
        input:       input{someArg: "foo"},
        expected:    expected{someReturnArg: nil},
    },
}

for i, tc := range testCases {
	t.Run(fmt.Sprintf("%d: %s", i, tc.description), func(t *testing.T) {
		res := someFunc(tc.input)
		if res != tc.expected {
			t.Errorf("expected: %v; got: %v", tc.expected, res)
		}
	})
}
```

Take the tests you wrote in the previous exercise, and convert them to a table driven test. To see if your tests work,
tests can be run as follows.

```
go test -v .
```

When you are done, validate that you have implemented the solution correctly by running the following command.

```
task test-solution
```

## AI

* Show me an example of table driven tests in Go, the programming language.