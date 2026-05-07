# Exercise

Sometimes code has one or many dependencies. One way to test code with dependencies is to set up the code under test with a real version of all dependencies. But that's not always what you want. For example, imagine code that connects to a database to fetch data. To be able to test this, you would have to run a database locally, put data in it, and then assert that it all works as expected. This _is_ possible, and it's a form of integration testing. Writing and maintaining these tests typically takes more time takes more time, and the tests are often slower to run, which is why we sometimes choose to pass a "fake" database as a dependency to our code - one that we can make behave exactly as we want. That fake database is called a "mock".

In fact, before we start mixing up terminology, these are the common types of "fake" implementations used in testing:

* **Stub:** A stub provides predefined responses to calls, so the code under test gets the data it needs without using the real dependency. It usually doesn’t contain real logic and you typically don’t assert on how it was called.
* **Fake:** A fake has a working, but simplified or in‑memory implementation (e.g., an in‑memory database). It behaves more like the real thing than a stub, but is still only suitable for tests.
* **Mock:** A mock is a programmable object that both simulates behavior and records how it was used. You typically set expectations on it (which methods you expect to be called with which arguments) and assert those expectations in the test.

So, a mock is an object that doubles as the real thing. In order for that to work, it conforms to the same interface that the real thing conforms to. That means that the only way to use mocks, is to use interfaces, because else we wouldn't be able to pass the fake thing as a replacement of the real.

Because the mock is based on an interface, we can use a code generator to generate the mock implementation. A mock implementation usually allows you to do something like:

```
someMock := NewMockOfSomeInterface()
someMock.Expect(someMethod).ToBeCalled(once).WithArgs(someArgs...).AndReturn(someValue)
```

Testify comes with some mocking functionality out of the box, but it still requires you to write your own mocks. In this assignment we'll generate our mock code to speed things up a little, using [Mockery](https://vektra.github.io/mockery/latest/), which actually uses Testify's libraries under the hood. Install it as follows.

```shell
go get -tool -modfile=../../../tools.mod github.com/vektra/mockery/v3
```

Since it's already part of the tools.mod files, as in the previous exercises, this command will probably not change any files.

We then generate mocks using the following command. 

```shell
task main:gen
```

This command runs the `gen` task defined in the Taskfile at the root of this project, which is:

```shell
go generate $(go list ./...)
```

This, in turn, runs the `go generate` command on all the Go files in this project, which will find generate.go in the current directory. Generate.go calls the tool - Mockery - installed using `go get -tool`. Mockery looks for a configuration file called .mockery.yml, which also lives in the current directory.

Then it's time for the actual exercise. Create a test for the code in create_user_use_case.go. You will need to use the mock object in ./mocks for one of the test cases. Example: 

```go
userRepo := mocks.NewMockUserRepository(t)

// Set up an expectation. Just start typing and let your autocomplete suggest what to use.
userRepo.EXPECT().SomeMethod("someArg").Return(nil).Once()
``` 

To see if your tests work, tests can be run as follows.

```shell
go test -v .
```

When you are done, validate that you have implemented the solution correctly by running the following command.

```shell
task test-solution
```

## AI Prompts

* Explain to me the different types of testing typically used in software engineering, for example unit tests, integration tests, system tests, regression tests.
* Explain to me what type of testing to use when in software engineering.