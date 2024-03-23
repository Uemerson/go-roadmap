# About

Goroutines allow us to write concurrent programs in Go. Things like web servers handling thousands of requests or a website rendering new pages while also concurrently making network requests are a few example of concurrency.

In Go, each of these concurrent tasks are called Goroutines.

# Goroutines

A goroutine is a lightweight thread managed by the Go runtime.

```
go f(x, y, z)
```

starts a new goroutine running

```
f(x, y, z)
```

The evaluation of `f, x, y`, and `z` happens in the current goroutine and the execution of `f` happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized. The `sync` package provides useful primitives, although you won't need them much in Go as there are other primitives.

[See this example.](./examples/goroutines/main.go)

# More of goroutines

They're called goroutines because the existing terms—threads, coroutines, processes, and so on—convey inaccurate connotations. A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap, and grow by allocating (and freeing) heap storage as required.

Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run. Their design hides many of the complexities of thread creation and management.

Prefix a function or method call with the go keyword to run the call in a new goroutine. When the call completes, the goroutine exits, silently. (The effect is similar to the Unix shell's & notation for running a command in the background.)

```
go list.Sort()  // run list.Sort concurrently; don't wait for it.
```

A function literal can be handy in a goroutine invocation.

```
func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // Note the parentheses - must call the function.
}
```

In Go, function literals are closures: the implementation makes sure the variables referred to by the function survive as long as they are active.

These examples aren't too practical because the functions have no way of signaling completion. For that, we need channels.

# Advantages of Goroutines
- Goroutines are cheaper than threads.
- Goroutine are stored in the stack and the size of the stack can grow and shrink according to the requirement of the program. But in threads, the size of the stack is fixed.
- Goroutines can communicate using the channel and these channels are specially designed to prevent race conditions when accessing shared memory using Goroutines.
- Suppose a program has one thread, and that thread has many Goroutines associated with it. If any of Goroutine blocks the thread due to resource requirement then all the remaining Goroutines will assign to a newly created OS thread. All these details are hidden from the programmers.

# Anonymous Goroutine
In Go language, you can also start Goroutine for an anonymous function or in other words, you can create an anonymous Goroutine simply by using go keyword as a prefix of that function as shown in the below Syntax:

```
// Anonymous function call
go func (parameter_list){
// statement
}(arguments)
```

[See this example.](./examples/anonymous-goroutine/main.go)

# WaitGroups

Wait groups are a mechanism for managing concurrency in Go. They provide a way to wait for a group of goroutines to complete their execution before proceeding further.

Wait groups are created using the sync package, and they provide three essential methods: `Add()`, `Done()`, and `Wait()`.

- `Add()` is used to add the number of goroutines that need to be waited upon.
- `Done()` is called by each goroutine when it finishes its work, decrementing the internal counter of the wait group.
- `Wait()` is used to block the execution of the goroutine until all the goroutines have called `Done()`.

By incrementing the wait group counter with `Add()`, each goroutine signals that it needs to be waited upon. When a goroutine finishes its work, it calls `Done()`, reducing the wait group counter. The main goroutine or another goroutine can then call `Wait()` to block until the wait group counter reaches zero.

[In this example](./examples/wait-group/main.go), each worker performs some work, and before starting the work, we increment the wait group counter using wg.Add(1). After each worker finishes its work, it calls wg.Done(), decrementing the wait group counter. Finally, the main goroutine calls wg.Wait() to block until all the workers have completed, and then it proceeds to print "All workers finished".

# Reference(s)

[A Tour of Go](https://go.dev/tour/concurrency/1)

[Effective Go](https://go.dev/doc/effective_go#goroutines)

[Goroutines – Concurrency in Golang](https://www.geeksforgeeks.org/goroutines-concurrency-in-golang/)

[Concurrency in Go: Channels and WaitGroups](https://medium.com/goturkiye/concurrency-in-go-channels-and-waitgroups-25dd43064d1)