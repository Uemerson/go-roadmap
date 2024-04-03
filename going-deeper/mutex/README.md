# Mutex
Go allows us to run code concurrently using goroutines. However, when concurrent processes access the same piece of data, it can lead to race conditions. Mutexes are data structures provided by the sync package. They can help us place a lock on different sections of data so that only one goroutine can access it at a time.

A mutex, short for mutual exclusion, is used to protect shared resources from simultaneous access by multiple goroutines. It ensures that only one goroutine can access a critical section of code at a time. Race conditions occur when multiple goroutines access and modify shared data concurrently, leading to unpredictable and erroneous behavior. Mutexes prevent race conditions by allowing only one goroutine to acquire the lock and access the shared resource, while other goroutines wait until the lock is released.

Mutexes are data structures provided by the standard `sync` package.

# How to prevent race conditions

In Go, the sync package provides the Mutex type, which includes two main methods: `Lock()` and `Unlock()`.

To understand how a mutex solves race conditions, let’s consider [this example](./examples/race-condition/main.go) without using a mutex.

We need to use the `sync.Mutex` type to prevent multiple goroutines from accessing counter at the same time. [See this example](./examples/fix-race-condition-with-mutex/main.go)

# Where Not to Use Mutex
1. High Contention: If many goroutines are frequently trying for the same lock, the performance of mutexes can degrade. In such cases, consider using alternative synchronization primitives like sync.RWMutex or channel-based communication patterns.
2. Deadlock Risks: Improper use of mutexes can lead to deadlocks, where goroutines end up waiting indefinitely for a lock to be released. Avoid complex nesting of locks or forgetting to unlock the mutex.

# Reference(s)

[Understanding Mutex in Go](https://kamnagarg-10157.medium.com/understanding-mutex-in-go-5f41199085b9)