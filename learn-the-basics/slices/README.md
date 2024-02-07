# About

Slices are similar to arrays but are more powerful and flexible. Like arrays, slices are also used to store multiple values of the same type in a single variable. However, unlike arrays, the length of a slice can grow and shrink as you see fit.

# Slices

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```
a[low : high]
```

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

```
a[1:4]
```

[See this example.](./examples/slices/main.go)

# Slices are like references to arrays

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

[See this example.](./examples/slices-pointers/main.go)

# Slice literals

A slice literal is like an array literal without the length.

This is an array literal:

```
[3]bool{true, true, false}
```

And this creates the same array as above, then builds a slice that references it:

```
[]bool{true, true, false}
```

[See this example.](./examples/slices-literals/main.go)

# Slice defaults

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

```
var a [10]int
```

these slice expressions are equivalent:

```
a[0:10]
a[:10]
a[0:]
a[:]
```

[See this example.](./examples/slices-bounds/main.go)

# Slice length and capacity

A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

[See this example.](./examples/slice-len-cap/main.go)

# Nil slices

The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.

[See this example.](./examples/nil-slices/main.go)

# Creating a slice with make

Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

```
a := make([]int, 5) // len(a)=5
```

To specify a capacity, pass a third argument to make:

```
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:] // len(b)=4, cap(b)=4
```

[See this example.](./examples/making-slices/main.go)

# Slices of slices

Slices can contain any type, including other slices.

[See this example.](./examples/slices-of-slice/main.go)

# Appending to a slice

It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.

```
func append(s []T, vs ...T) []T
```

The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

[See this example.](./examples/append/main.go)

# Reference(s)

[A Tour of Go](https://go.dev/tour/moretypes/7)
