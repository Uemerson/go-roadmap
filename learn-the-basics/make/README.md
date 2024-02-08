# Make

Golang’s built-in function make, helps us create and initialize slices, maps and channels, depending on the arguments that are provided to the function.

# Allocation with make

Back to allocation. The built-in function make(T, args) serves a purpose different from new(T). It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not \*T). The reason for the distinction is that these three types represent, under the covers, references to data structures that must be initialized before use. A slice, for example, is a three-item descriptor containing a pointer to the data (inside an array), the length, and the capacity, and until those items are initialized, the slice is nil. For slices, maps, and channels, make initializes the internal data structure and prepares the value for use. For instance,

```
make([]int, 10, 100)
```

allocates an array of 100 ints and then creates a slice structure with length 10 and a capacity of 100 pointing at the first 10 elements of the array. (When making a slice, the capacity can be omitted; see the section on slices for more information.) In contrast, new([]int) returns a pointer to a newly allocated, zeroed slice structure, that is, a pointer to a nil slice value.

These examples illustrate the difference between new and make.

```
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)
```

Remember that make applies only to maps, slices and channels and does not return a pointer. To obtain an explicit pointer allocate with new or take the address of a variable explicitly.

# Create maps, slice, channels using Make function

To create a slice using make:

```
var intSlice = make([]int, 10)        // when length and capacity is same
var strSlice = make([]string, 10, 20) // when length and capacity is different
```

To create a map using make:

```
var employee = make(map[string]int)
```

To create a map using make:

```
channelName := make(chan int)
```

- channelName - name of the channel
- (chan int) - indicates that the channel is of integer type

# Reference(s)

[Effective Go](https://go.dev/doc/effective_go#allocation_make)

[How to create Slice using Make function in Golang?](https://www.golangprograms.com/how-to-create-slice-using-make-function-in-golang.html)

[How to create Map using the make function in Go?](https://www.golangprograms.com/golang-package-examples/how-to-create-map-using-the-make-function-in-go.html)

[Go Channel](https://www.programiz.com/golang/channel#channel)
