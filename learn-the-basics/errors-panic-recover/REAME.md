# Errors/Panic/Recover

In lieu of adding exception handlers, the Go creators exploited Go’s ability to return multiple values. The most commonly used Go technique for issuing errors is to return the error as the last value in a return.

A panic typically means something went unexpectedly wrong. Mostly used to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.

Panic recovery in Go depends on a feature of the language called deferred functions. Go has the ability to guarantee the execution of a function at the moment its parent function returns. This happens regardless of whether the reason for the parent function’s return is a return statement, the end of the function block, or a panic.

# Error

Error handling in Go is a little different than other mainstream programming languages like Java, JavaScript, or Python. Go’s built-in errors don’t contain stack traces, nor do they support conventional try/catch methods to handle them. Instead, errors in Go are just values returned by functions, and they can be treated in much the same way as any other datatype - leading to a surprisingly lightweight and simple design.

[See this example.](./examples/error/main.go)

# The Error Type

The error type in Go is implemented as the following interface:

```
type error interface {
    Error() string
}
```

So basically, an error is anything that implements the Error() method, which returns an error message as a string. It’s that simple!

# Constructing Errors

Errors can be constructed on the fly using Go’s built-in errors or fmt packages.

[This example has a function that uses the errors package to return a new error with a static error message.](./examples/errors/main.go)

[This example use the fmt package can be used to add dynamic data to the error, such as an int, string, or another error.](./examples/error/main.go)

- Errors can be returned as nil, and in fact, it’s the default, or “zero”, value of on error in Go. This is important since checking if err != nil is the idiomatic way to determine if an error was encountered (replacing the try/catch statements you may be familiar with in other programming languages).

- Errors are typically returned as the last argument in a function. Hence in our example above, we return an int and an error, in that order.

- When we do return an error, the other arguments returned by the function are typically returned as their default “zero” value. A user of a function may expect that if a non-nil error is returned, then the other arguments returned are not relevant.

- Lastly, error messages are usually written in lower-case and don’t end in punctuation. Exceptions can be made though, for example when including a proper noun, a function name that begins with a capital letter, etc.

# Defining Expected Errors

Another important technique in Go is defining expected Errors so they can be checked for explicitly in other parts of the code. This becomes useful when you need to execute a different branch of code if a certain kind of error is encountered.

## Defining Sentinel Errors

Building on the Divide function from earlier, we can improve the error signaling by pre-defining a “Sentinel” error. Calling functions can explicitly check for this error using errors.Is

    A sentinel error in Go is an error that is predefined as a global variable. It’s typically used to represent an “expected” error

[See this example.](./examples/errors-is/main.go)

## Defining Custom Error Types

Many error-handling use cases can be covered using the strategy above, however, there can be times when you might want a little more functionality. Perhaps you want an error to carry additional data fields, or maybe the error’s message should populate itself with dynamic values when it’s printed.

You can do that in Go by implementing custom errors type.

[In this example](./examples/custom-error-types/main.go) is a slight rework of the previous example. Notice the new type DivisionError, which implements the Error interface. We can make use of errors.As to check and convert from a standard error to our more specific DivisionError.

## Customizing error tests with Is and As methods

The errors.Is function examines each error in a chain for a match with a target value. By default, an error matches the target if the two are equal. In addition, an error in the chain may declare that it matches a target by implementing an Is method.

[As an example](./examples/custom-error-is/main.go), consider this error inspired by the Upspin error package which compares an error against a template, considering only fields which are non-zero in the template:

# Reference(s)

[Effective Error Handling in Golang](https://earthly.dev/blog/golang-errors/)

[Writing clean code in go: sentinel errors](https://medium.com/gopher-time/writing-clean-code-in-go-sentinel-errors-5ad93a30bc8e)

[Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
