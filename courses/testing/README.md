# Testing

## Why testing

Testing helps you catch bugs early, gives you confidence when making changes, and serves as documentation for how your
code should behave. A good test suite means you can refactor without fear and ship code knowing it actually works.

In Go, testing is built into the language with the `testing` package. Tests live alongside your code in `*_test.go`
files and run with `go test`.

## Test suite

A test suite is a collection of test cases that are run together. In general software terms, it's any group of tests
that
verify related functionality or the entire application.

In Go, a test suite typically refers to all the `*_test.go` files in a package or directory. When you run `go test`, it
executes all tests in the current package—that's your test suite. You can also run tests across multiple packages with
`go test ./...` to run your entire project's test suite.

## Types of testing

- **Unit tests** test individual functions or methods in isolation. They're fast and help you verify that each piece
  works correctly on its own.
- **Integration tests** test how multiple components work together. They're slower but catch issues that unit tests
  miss, like database interactions or API calls.
- **End-to-end tests** test the entire system from a user's perspective. They're the slowest but give the most
  confidence that everything works together.

Start with unit tests for your core logic, add integration tests for critical paths, and use end-to-end tests sparingly
for the most important user journeys.

## Test coverage

Test coverage measures what percentage of your code is executed by tests. It's a useful metric but not the whole
story—100% coverage doesn't mean your tests are good, just that they run all the code.

Aim for high coverage on critical business logic and low coverage on trivial code. A good target is 70-80% for most
projects. Focus on testing behavior and edge cases, not just executing lines of code.

## Mocks

Mocks are test doubles that replace real dependencies (like databases or external APIs) with controlled versions. They
let you test your code in isolation without needing the actual infrastructure running.

### Stub, mock, spy, fake

These terms are often used interchangeably, but technically:

- **Stubs** return hardcoded responses.
- **Mocks** verify that specific methods were called with expected arguments.
- **Spies** record how they were called so you can inspect it later.
- **Fakes** are simplified working implementations (like an in-memory database)

In practice, most people just say "mock" for all of these. What matters is isolating the code you're testing.

## Test Driven Development

Test Driven Development (TDD) means writing tests before writing the implementation. The cycle is: write a failing test,
write the minimal code to make it pass, then refactor.

TDD forces you to think about how your code will be used before you write it. This often leads to better APIs and more
testable code. It also ensures you actually write tests instead of putting them off.

## AI Prompts

* Give me an overview of the five most uses types of testing in software applications and explain which to use when.
* Explain the trade-offs between 100% test coverage in software applications versus not reaching 100%.
* Convince me that I do not need to write a unit test for every piece of code that I write.
* Convince me that I should write a unit test for every piece of code that I write.

**Disclaimer:** This file was partially written using AI.