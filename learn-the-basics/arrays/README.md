# About

In Go an array is a collection of elements of the same type with a fixed size defined when the array is created.

# Arrays

The type [n]T is an array of n values of type T.

The expression

```
var a [10]int
```

declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

[See this example.](./examples/arrays/main.go)

# Reference(s)

[A Tour of Go](https://go.dev/tour/moretypes/6)
