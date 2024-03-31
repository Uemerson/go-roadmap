# About

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`

# Channels

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

[In this example](./examples/channels/main.go) we create a new channel with make(chan val-type). `Channels are typed by the values they convey.` Send a value into a channel using the channel <- syntax. [In this example](./examples/channels/main.go) we send "ping" to the messages channel we made above, from a new goroutine. 
When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.

By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.

# Sending and receiving from a channel

The syntax to send and receive data from a channel is given below,

```
data := <- a // read from channel a
a <- data // write to channel a
```

The direction of the arrow with respect to the channel specifies whether the data is sent or received.

In the first line, the arrow points outwards from `a` and hence we are reading from channel `a` and storing the value to the variable `data`.

In the second line, the arrow points towards a and hence we are writing to channel `a`.

# Channel Buffering

By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.

[In this example](./examples/channel-buffering/main.go) we make a channel of strings buffering up to 2 values.

Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive. Later we can receive these two values as usual.

# Channel Synchronization

We can use channels to synchronize execution across goroutines. [Here’s an example](./examples/channel-synchronization/main.go) of using a blocking receive to wait for a goroutine to finish. When waiting for multiple goroutines to finish, you may prefer to use a `WaitGroup.` The `done` channel will be used to notify another goroutine that this function’s work is done. `<-done` will block the program execution until we receive a notification from the worker on the channel.

# Channel Directions

When using channels as function parameters, you can specify if a channel is meant to only send or receive values. This specificity increases the type-safety of the program.

[In this example](./examples/channel-directions/main.go) the ping function only accepts a channel for sending values. It would be a compile-time error to try to receive on this channel. The pong function accepts one channel for receives (pings) and a second for sends (pongs).

# Select

Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.

[In this example](./examples/channel-select/main.go) each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines. We’ll use select to await both of these values simultaneously, printing each one as it arrives.

# Non-Blocking Channel Operations

Basic sends and receives on channels are blocking. However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.

[In this example](./examples/non-blocking-channel-operations/main.go) there’s a non-blocking receive. If a value is available on messages then select will take the <-messages case with that value. If not it will immediately take the default case. 

A non-blocking send works similarly. In the example `msg` cannot be sent to the messages channel, because the channel has no buffer and there is no receiver. Therefore the default case is selected.

We can use multiple `cases` above the default clause to implement a multi-way non-blocking select. Here we attempt non-blocking receives on both messages and signals.

# Closing Channels

Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers.

[In this example](./examples/closing-channels/main.go) we’ll use a jobs channel to communicate work to be done from the `main()` goroutine to a worker goroutine. When we have no more jobs for the worker we’ll close the jobs channel.

# Range over Channels

 We can use range syntax to iterate over values received from a channel.

[See this example](./examples/range-over-channels/main.go)

# Deadlock

One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.

Similarly, if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write data on that channel, else the program will panic.

[In this example](./examples/deadlock-channels/main.go), a channel `ch` is created and we send `5` to the channel in line no. 6 `ch <- 5`. In this program no other Goroutine is receiving data from the channel `ch`. Hence this program will panic with the following runtime error.

# Reference(s)

[Go by Example: Channels](https://gobyexample.com/channels)

[Channels](https://golangbot.com/channels/)