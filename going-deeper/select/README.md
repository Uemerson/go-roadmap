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

# Combining goroutines and channels

Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.

[In this example](./examples/select-across-two-channels/main.go) we’ll select across two channels.

Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.

We’ll use `select` to await both of these values simultaneously, printing each one as it arrives.

We receive the values "one" and then "two" as expected.

Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.

# Reference(s)

[A Tour of Go](https://go.dev/tour/concurrency/5)

[Go by Example: Select](https://gobyexample.com/select)
