# Exercise

Another popular testing framework is [GoConvey](https://github.com/smartystreets/goconvey). This framework is more
similar to RSpec. Install it using:

```shell
go get github.com/smartystreets/goconvey
```

As before, since GoConvey is already being used in this project, installation probably did nothing, and again that's 
fine.

A typical GoConvey test looks like this:

```go
Convey("If x is true", t, func() {
    x := true
    
    Convey("If y is true", func() {
        y := true
        
        Convey("Returns z", func() {
            So(someFunc(x, y), ShouldEqual, "z")
        })
    })

    Convey("If y is false", func() {
        y := false
        
        Convey("Returns error", func() {
            So(someFunc(x, y), ShouldEqual, someErr)
        })
    })
})
```

In the example above, the following happens. First the root `Convey()` call receives the reference to `*testing.T` and
is executed. Then the first child `Convey()` call. Then the first child of the first child. This keeps on going until
there is no other first child. Then the process starts over from the root and executes the first child, the first child,
etc., and of the last nested `Convey()` call, it then executes the second child `Convey()` call. This way the state is
reset and rebuilt for every individual leaf case. For more info,
read [the documentation](https://github.com/smartystreets/goconvey/wiki/Execution-order) on execution order.

As in the previous exercise, write tests for the implementations in calculator.go, instead this time use `Convey()` 
instead of testify.

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