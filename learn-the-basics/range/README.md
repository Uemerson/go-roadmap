# About

Range is used with For Loops to iterate over each element in arrays, strings and other data structures.

# Range

The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

[See this example.](./examples/range/main.go)

# Range continued

You can skip the index or value by assigning to \_.

```
for i, _ := range pow
for _, value := range pow
```

If you only want the index, you can omit the second variable.

```
for i := range pow
```

[See this example.](./examples/range-continued/main.go)

# Reference(s)
