# Errors/Panic/Recover

In lieu of adding exception handlers, the Go creators exploited Go’s ability to return multiple values. The most commonly used Go technique for issuing errors is to return the error as the last value in a return.

A panic typically means something went unexpectedly wrong. Mostly used to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.

Panic recovery in Go depends on a feature of the language called deferred functions. Go has the ability to guarantee the execution of a function at the moment its parent function returns. This happens regardless of whether the reason for the parent function’s return is a return statement, the end of the function block, or a panic.

# Error

[This markdown contains details on how errors work in Go.](./ERROR.md)

# Defer

[This markdown contains details on how defer work in Go.](./DEFER.md)

# Panic

Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

# Recover

Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

[Here’s an example program](./examples/mechanics-panic-defer/main.go) that demonstrates the mechanics of panic and defer.

- The function g takes the int i, and panics if i is greater than 3, or else it calls itself with the argument i+1. The function f defers a function that calls recover and prints the recovered value (if it is non-nil). Try to picture what the output of this program might be before reading on.

- **If we remove the deferred function from f the panic is not recovered and reaches the top of the goroutine’s call stack, terminating the program.**

The convention in the Go libraries is that even when a package uses panic internally, its external API still presents explicit error return values.

For a real-world example of panic and recover, see the [json package](https://pkg.go.dev/encoding/json) from the Go standard library. It encodes an interface with a set of recursive functions. If an error occurs when traversing the value, panic is called to unwind the stack to the top-level function call, which recovers from the panic and returns an appropriate error value (see the ’error’ and ‘marshal’ methods of the encodeState type in [encode.go](https://go.dev/src/encoding/json/encode.go)).

The convention in the Go libraries is that even when a package uses panic internally, its external API still presents explicit error return values.

# Reference(s)

[Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
