# Context

The context package provides a standard way to solve the problem of managing the state during a request. The package satisfies the need for request-scoped data and provides a standardized way to handle: Deadlines, Cancellation Signals, etc.

# Creating a Context

To create a context, you can use the `context.Background()` function, which returns an empty, non-cancelable context as the root of the context tree. You can also create a context with a specific timeout or deadline using `context.WithTimeout()` or `context.WithDeadline()` functions.

# Creating a Context with Timeout

[In this example](./examples/context-with-timeout/main.go), we create a context with a timeout of 2 seconds and use it to simulate a time-consuming operation

```
In this example, the performTask function simulates a long-running task that takes 5 seconds to complete. However, since the context has a timeout of only 2 seconds, the operation is terminated prematurely, resulting in a timeout.
```

# Reference(s)

[The Complete Guide to Context in Golang: Efficient Concurrency Management](https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea)
