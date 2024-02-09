# About

Structs are user-defined types that help us create a collection of data describing a single entity.

# Structs

A struct is a collection of fields.

[See this example.](./examples/structs/main.go)

# Struct Fields

Struct fields are accessed using a dot.

[See this example.](./examples/struct-fields/main.go)

# Pointers to structs

Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (\*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

[See this example.](./examples/struct-pointers/main.go)

# Struct Literals

A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value.

[See this example.](./examples/struct-literals/main.go)

# Reference(s)

[A Tour of Go](https://go.dev/tour/moretypes/2)
