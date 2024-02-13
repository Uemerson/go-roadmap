# Select

The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready. The select statement is just like switch statement, but in the select statement, case statement refers to communication, i.e. sent or receive operation on the channel.

[See this example.](./examples/select/main.go)

# Default Selection

The default case in a select is run if no other case is ready.

Use a default case to try a send or receive without blocking:

```
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

[See this example.](./examples/default-selection/main.go)

# Reference(s)

[A Tour of Go](https://go.dev/tour/concurrency/5)
