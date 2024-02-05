# Functions

A function is a group of statements that exist within a program for the purpose of performing a specific task. At a high level, a function takes an input and returns an output.

Function allows you to extract commonly used block of code into a single component.

**The single most popular Go function is main(), which is used in every independent Go program.**

# Naming Conventions for Golang Functions

- A name must begin with a letter, and can have any number of additional letters and numbers.
- A function name cannot start with a number.
- A function name cannot contain spaces.
- If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.
- If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
- function names are case-sensitive (car, Car and CAR are three different variables).

# A Function in Golang

A declaration begins with the func keyword, followed by the name you want the function to have, a pair of parentheses (), and then a block containing the function's code.

[See this example.](./examples/function/main.go)

# Multiple Return Values

Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example to return both result and error values from a function.

[See this example.](./examples/multiple-return-values/main.go)

# Variadic Functions

Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

[Here’s a function](./examples/variadic-functions/main.go) that will take an arbitrary number of ints as arguments.

# Anonymous Functions

An anonymous function is a function that was declared without any named identifier to refer to it. Anonymous functions can accept inputs and return outputs, just as standard functions do.

[Assigning function to the variable](./examples/assigning-function-to-variable/main.go)

[Passing arguments to anonymous functions](./examples/passing-arguments-to-anonymous-functions/main.go)

[Function defined to accept a parameter and return value.](./examples/function-accept-parameter-return-value/main.go)

# Closures Functions

Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function.

[See this example.](./examples/closures/main.go)
