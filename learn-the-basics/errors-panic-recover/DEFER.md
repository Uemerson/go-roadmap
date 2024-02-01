# Defer

Go has the usual mechanisms for control flow: if, for, switch, goto. It also has the go statement to run code in a separate goroutine. Here I’d like to discuss some of the less common ones: defer, panic, and recover.

A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.

[For example](./examples/defer-bug/main.go), let’s look at a function that opens two files and copies the contents of one file to the other.

This example works, but there is a bug. If the call to os.Create fails, the function will return without closing the source file. This can be easily remedied by putting a call to src.Close before the second return statement, but if the function were more complex the problem might not be so easily noticed and resolved. By introducing defer statements we can ensure that the files are always closed, [like this example.](./examples/defer/main.go)

Defer statements allow us to think about closing each file right after opening it, guaranteeing that, regardless of the number of return statements in the function, the files will be closed.

## The behavior of defer statements is straightforward and predictable.

There are three simple rules:

1 - A deferred function’s arguments are evaluated when the defer statement is evaluated.

[In this example](./examples/defer-print-i/main.go), the expression “i” is evaluated when the Println call is deferred. The deferred call will print “0” after the function returns.

2 - Deferred function calls are executed in Last In First Out order after the surrounding function returns.

[This example](./examples/defer-prints-3210/main.go) prints “3210”

3 - Deferred functions may read and assign to the returning function’s named return values.

[In this example](./examples/defer-function-increments/main.go), a deferred function increments the return value i after the surrounding function returns. Thus, this function returns 2

# Reference(s)

[Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
