# Context

The context package provides a standard way to solve the problem of managing the state during a request. The package satisfies the need for request-scoped data and provides a standardized way to handle: Deadlines, Cancellation Signals, etc.

# What is Context?

Context is a built-in package in the Go standard library that provides a powerful toolset for managing concurrent operations. It enables the propagation of cancellation signals, deadlines, and values across goroutines, ensuring that related operations can gracefully terminate when necessary. With context, you can create a hierarchy of goroutines and pass important information down the chain.

## Managing Concurrent API Requests

Consider a scenario where you need to fetch data from multiple APIs concurrently. By using context, you can ensure that all the API requests are canceled if any of them exceeds a specified timeout.

[In this example](./examples/managing-concurrent-api-requests/main.go), we create a context with a timeout of 5 seconds. We then launch multiple goroutines to fetch data from different APIs concurrently. The http.NewRequestWithContext() function is used to create an HTTP request with the provided context. If any of the API requests exceed the timeout duration, the context's cancellation signal is propagated, canceling all other ongoing requests.

# Creating a Context

To create a context, you can use the `context.Background()` function, which returns an empty, non-cancelable context as the root of the context tree. You can also create a context with a specific timeout or deadline using `context.WithTimeout()` or `context.WithDeadline()` functions.

## Creating a Context with Timeout

[In this example](./examples/context-with-timeout/main.go), we create a context with a timeout of 2 seconds and use it to simulate a time-consuming operation

In this example, the `performTask` function simulates a long-running task that takes 5 seconds to complete. However, since the context has a timeout of only 2 seconds, the operation is terminated prematurely, resulting in a timeout.

# Propagating Context

Once you have a context, you can propagate it to downstream functions or goroutines by passing it as an argument. This allows related operations to share the same context and be aware of its cancellation or other values.

## Propagating Context to Goroutines

[In this example](./examples/propagating-context-goroutines/main.go), we create a parent context and propagate it to multiple goroutines to perform concurrent tasks.

In this example, we create a parent context using `context.Background()`. We then use `context.WithValue()` to attach a user ID to the context. The context is then passed to the `performTask` goroutine, which retrieves the user ID using `ctx.Value()`.

# Retrieving Values from Context

In addition to propagating context, you can also retrieve values stored within the context. This allows you to access important data or parameters within the scope of a specific goroutine or function.

## Retrieving User Information from Context

[In this example](./examples/retrieving-values-from-context/main.go), we create a context with user information and retrieve it in a downstream function.

In this example, we create a context using `context.WithValue()` and store the user ID. The context is then passed to the `processRequest` function, where we retrieve the user ID using type assertion and use it for further processing.

# Cancelling Context

Cancellation is an essential aspect of context management. It allows you to gracefully terminate operations and propagate cancellation signals to related goroutines. By canceling a context, you can avoid resource leaks and ensure the timely termination of concurrent operations.

## Cancelling Context

[In this example](./examples/cancelling-context/main.go), we create a context and cancel it to stop ongoing operations.

In this example, we create a context using `context.WithCancel()` and defer the cancellation function. The `performTask` goroutine continuously performs a task until the context is canceled. After 2 seconds, we call the cancel function to initiate the cancellation process. As a result, the goroutine detects the cancellation signal and terminates the task.

# Timeouts and Deadlines

Setting timeouts and deadlines is crucial when working with context in Golang. It ensures that operations complete within a specified timeframe and prevents potential bottlenecks or indefinite waits.

## Setting a Deadline for Context

[In this example](./examples/setting-a-deadline-for-context/main.go), we create a context with a deadline and perform a task that exceeds the deadline.

In this example, we create a context with a deadline of 2 seconds using `context.WithDeadline()`. The `performTask` goroutine waits for the context to be canceled or for the deadline to be exceeded. After 3 seconds, we let the program exit, triggering the deadline exceeded error.

# Context in HTTP Requests

Context plays a vital role in managing HTTP requests in Go. It allows you to control request cancellation, timeouts, and pass important values to downstream handlers.

## Using Context in HTTP Requests

[In this example](./examples/using-context-in-http-requests/main.go), we make an HTTP request with a custom context and handle timeouts.

In this example, we create a context with a timeout of 2 seconds using `context.WithTimeout()`. We then create an HTTP request with the custom context using `http.NewRequestWithContext()`. The context ensures that if the request takes longer than the specified timeout, it will be canceled.

# Context in Database Operations

Context is also useful when dealing with database operations in Golang. It allows you to manage query cancellations, timeouts, and pass relevant data within the database transactions.

## Using Context in Database Operations

[In this example](./examples/using-context-in-database-operations/main.go), we demonstrate how to use context with a PostgreSQL database operation.

In this example, we create a context with a timeout of 2 seconds using `context.WithTimeout()`. We then open a connection to a PostgreSQL database using the `sql.Open()` function. When executing the database query with `db.QueryContext()`, the context ensures that the operation will be canceled if it exceeds the specified timeout.

# Reference(s)

[The Complete Guide to Context in Golang: Efficient Concurrency Management](https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea)
