# Context

The context package provides a standard way to solve the problem of managing the state during a request. The package satisfies the need for request-scoped data and provides a standardized way to handle: Deadlines, Cancellation Signals, etc.

# Creating a Context

To create a context, you can use the `context.Background()` function, which returns an empty, non-cancelable context as the root of the context tree. You can also create a context with a specific timeout or deadline using `context.WithTimeout()` or `context.WithDeadline()` functions.

# Reference(s)

[The Complete Guide to Context in Golang: Efficient Concurrency Management](https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea)
