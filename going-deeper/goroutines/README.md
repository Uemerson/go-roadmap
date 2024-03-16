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

# Reference(s)

[A Tour of Go](https://go.dev/tour/concurrency/1)