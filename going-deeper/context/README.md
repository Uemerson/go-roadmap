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

# Best Practices for Using Context

When working with context in Golang, it is essential to follow some best practices to ensure efficient and reliable concurrency management.

## Implementing Best Practices for Context Usage

Here are some best practices to consider:

1. Pass Context Explicitly: Always pass the context as an explicit argument to functions or goroutines instead of using global variables. This makes it easier to manage the context’s lifecycle and prevents potential data races.

2. Use context.TODO(): If you are unsure which context to use in a particular scenario, consider using `context.TODO()`. However, make sure to replace it with the appropriate context later.

3. Avoid Using context.Background(): Instead of using `context.Background()` directly, create a specific context using `context.WithCancel()` or `context.WithTimeout()` to manage its lifecycle and avoid resource leaks.

4. Prefer Cancel Over Timeout: Use `context.WithCancel()` for cancellation when possible, as it allows you to explicitly trigger cancellation when needed. `context.WithTimeout()` is more suitable when you need an automatic cancellation mechanism.

5. Keep Context Size Small: Avoid storing large or unnecessary data in the context. Only include the data required for the specific operation.

6. Avoid Chaining Contexts: Chaining contexts can lead to confusion and make it challenging to manage the context hierarchy. Instead, propagate a single context throughout the application.

7. Be Mindful of Goroutine Leaks: Always ensure that goroutines associated with a context are properly closed or terminated to avoid goroutine leaks.

# Context in Real-World Scenarios

Context in Golang is widely used in various real-world scenarios. Let’s explore some practical examples where context plays a crucial role.

## Context in Microservices

In a microservices architecture, each service often relies on various external dependencies and communicates with other services. Context can be used to propagate important information, such as authentication tokens, request metadata, or tracing identifiers, throughout the service interactions.

## Context in Web Servers

Web servers handle multiple concurrent requests, and context helps manage the lifecycle of each request. Context can be used to set timeouts, propagate cancellation signals, and pass request-specific values to the different layers of a web server application.

## Context in Test Suites

When writing test suites, context can be utilized to manage test timeouts, control test-specific configurations, and enable graceful termination of tests. Context allows tests to be canceled or skipped based on certain conditions, enhancing test control and flexibility.

# Common Pitfalls to Avoid

1. Not propagating the context — Child functions need the context passed to them in order to honor cancelation. Don’t create contexts and keep them confined to one function.

2. Forgetting to call cancel — When done with a cancelable context, call the cancel function. This releases resources and stops any associated goroutines.

3. Leaking goroutines — Goroutines started with a context must check the Done channel to exit properly. Otherwise they risk leaking when the context is canceled.

4. Using basic context.Background for everything — Background lacks cancelation and timeouts. Use the WithCancel, WithTimeout, or WithDeadline functions to add control.

5. Passing nil contexts — Passing nil instead of a real context causes panics. Make sure context is never nil when passing it.

6. Checking context too early — Don’t check context conditions like Done() early in an operation. This risks canceling before the work starts.

7. Using blocking calls — Blocking calls like file/network IO should be wrapped to check context cancellation. This avoids hanging.

8. Overusing contexts — Contexts are best for request-scoped operations. For globally shared resources, traditional patterns may be better.

9. Assuming contexts have timeouts — The context.Background has no deadline. Add timeouts explicitly when needed.

10. Forgetting contexts expire — Don’t start goroutines with a context and assume they will run forever. The context may expire.

# Context and Goroutine Leaks

Contexts in Go are used to manage the lifecycle and cancellation signaling of goroutines and other operations. A root context is usually created, and child contexts can be derived from it. Child contexts inherit cancellation from their parent contexts.

If a goroutine is started with a context, but does not properly exit when that context is canceled, it can result in a goroutine leak. The goroutine will persist even though the operation it was handling has been canceled.

[Here is an example](./examples/context-and-goroutine-leaks/main.go) of a goroutine leak due to improper context handling

In this example, the goroutine started with the context does not properly exit when that context is canceled. This will result in a goroutine leak, even though the main context is canceled.

To fix it, the goroutine needs to call the context’s Done() channel when the main context is canceled:

```
ctx, cancel := context.WithCancel(context.Background())
```

Now the goroutine will cleanly exit when the parent context is canceled, avoiding the leak. Proper context propagation and lifetime management is key to preventing goroutine leaks in Go programs.

# Using Context with Third-Party Libraries

Sometimes we need using third-party packages. However, many third-party libraries and APIs do not natively support context. So when using such libraries, you need to take some additional steps to integrate context usage properly:

- Wrap the third-party APIs you call in functions that accept a context parameter.
- In the wrapper function, call the third-party API as normal.
- Before calling the API, check if the context is done and return immediately if so. This propagates cancellation.
- After calling the API, check if the context is done and return immediately if so. This provides early return on cancellation.
- Make sure to call the API in a goroutine if it involves long-running work that could block.
- Define reasonable defaults for context values like timeout and deadline, so the API call isn’t open-ended.

[This example](./examples/context-with-third-party/main.go) provides context integration even with APIs that don’t natively support it. The key points are wrapping API calls, propagating cancellation, using goroutines, and setting reasonable defaults.

# Context(new features added in go1.21.0)

## func AfterFunc

Managing cleanup and finalization tasks is an important consideration in Go, especially when dealing with concurrency. The context package provides a useful tool for this through its `AfterFunc` function.

`AfterFunc` allows you to schedule functions to run asynchronously after a context ends. This enables deferred cleanup routines that will execute reliably once some operation is complete.

For example, imagine we have an API server that needs to process incoming requests from a queue. We spawn goroutines to handle each request:

```
func handleRequests(ctx context.Context) {
  for {
    req := queue.Get()
    go process(req)

    if ctx.Done() {
      break
    }
  }
}
```

But we also want to make sure any pending requests are processed if handleRequests has to exit unexpectedly. This is where `AfterFunc` can help.

We can schedule a cleanup function to run after the context is cancelled:

```
ctx, cancel := context.WithCancel(context.Background())

stop := context.AfterFunc(ctx, func() {
  // Process remaining queue
})

go handleRequests(ctx)

// Later when done...
cancel()
stop() // Prevent cleanup
```

Now our cleanup logic will run after the context ends. But since we call stop(), it is canceled before executing.

`AfterFunc` allows deferred execution tied to a context’s lifetime. This provides a robust way to build asynchronous applications with proper finalization.

## func WithDeadlineCause

When using contexts with deadlines in Go, timeout errors are common — a context will routinely expire if an operation takes too long. But the generic “context deadline exceeded” error lacks detail on the source of the timeout.

This is where `WithDeadlineCause` comes in handy. It allows you to associate a custom error cause with a context’s deadline:

```
ctx, cancel := context.WithDeadlineCause(ctx, time.Now().Add(100*time.Millisecond),
           errors.New("RPC timeout"))
defer cancel()

// Simulate work
time.Sleep(200 * time.Millisecond)

// Print the error cause
fmt.Println(ctx.Err()) // prints "context deadline exceeded: RPC timeout"
```

Now if the deadline is exceeded, the context’s Err() method will return:

“context deadline exceeded: RPC timeout”

This extra cause string gives critical context on the source of the timeout. Maybe it was due to a backend RPC call failing, or a network request timing out.

Without the cause, debugging the timeout requires piecing together where it came from based on call stacks and logs. But `WithDeadlineCause` allows directly propagating the source of the timeout through the context.

Timeouts tend to cascade through systems — a low-level timeout bubbles up to eventually become an HTTP 500. Maintaining visibility into the original cause is crucial for diagnosing these issues.

`WithDeadlineCause` enables this by letting you customize the deadline exceeded error with contextual details. The error can then be inspected at any level of the stack to understand the timeout source.

# func withTimeoutCause

Managing timeouts is an important aspect of writing reliable Go programs. When using context timeouts, the error “context deadline exceeded” is generic and lacks detail on the source of the timeout.

The `WithTimeoutCause` function addresses this by allowing you to associate a custom error cause with a context’s timeout duration:

```
ctx, cancel := context.WithTimeoutCause(ctx, 100*time.Millisecond,
          errors.New("Backend RPC timed out"))
```

Now if that ctx hits the timeout deadline, the context’s Err() will return:

“context deadline exceeded: Backend RPC timed out”

This provides critical visibility into the source of the timeout when it propagates up a call stack. Maybe it was caused by a slow database query, or a backend RPC service timing out.

Without a customized cause, debugging timeouts requires piecing together logs and traces to determine where it originated. But `WithTimeoutCause` allows directly encoding the source of the timeout into the context error.

Some key benefits of using `WithTimeoutCause`:

- Improved debugging of cascading timeout failures
- Greater visibility into timeout sources as errors propagate
- More context for handling and recovering from timeout errors

`WithTimeoutCause` gives more control over timeout errors to better handle them programmatically and debug them when issues arise.

# Reference(s)

[The Complete Guide to Context in Golang: Efficient Concurrency Management](https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea)
