# Task

Open channels.go and implement the three functions you see there according to the function documentation (which is
written over the function).

Validate that you have implemented the solution correctly by running the tests:

```
go test -v .
```

Where `-v` means "verbose" and `.` is the current directory.

Channels work as follows:

```go
// Create a channel.
ch := make(chan int)

// Send on channel.
// This however will panic because there is no one reading from the channel.
ch <- 1

// Instead, we can spin up a goroutine that sends the value into the channel.
go func() {
    // This line will block until someone starts reading the channel, but since it is in a goroutine, which is sort
	// of a separate process, the program will not panic and the main program will not be blocked - only this goroutine
	// will be blocked, until the channel is read from. Then the goroutine function finishes and the goroutine is
	// cleaned up.
    ch <- 1
}()

// Read from the channel. This will block until a value is sent, which is done in the goroutine we've just started,
// so the program won't panic.
val := <-ch

// Prints "val = 1".
fmt.Println("val =", val)
```

Multiple values can be sent on a channel. You can use `for range` to read from a channel until the channel is closed.

```go
// Loop will automatically terminate once `close(someChannel)` happens.
for val := range someChannel {
	// Do something with val.
}
```

See also the [example](https://gobyexample.com/channels) in Go By Example.

## Debugging

There are several ways of debugging, but the easiest for now would be to print to stdout. It can be done as follows:

```go
// Print any type of value and end with a newline.
fmt.Println("val", 1, someVal)

// "Regular" printf. %v means "any type of value".
fmt.Printf("val: %v\n", someVal)
```

## AI

* Explain the basics of how channels work in Go, the programming language, with some examples.