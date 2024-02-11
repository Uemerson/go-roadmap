# About

An interface in Go, is a type that defines a set of methods. If we have a type (e.g. struct) that implements that set of methods, then we have a type that implements this interface.

# Interfaces

An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

[See this example.](./examples/interfaces/main.go)

Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on \*Vertex (the pointer type).

# Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

[See this example.](./examples/interfaces-are-satisfied-implicitly/main.go)

# Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

```
(value, type)
```

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.

[See this example.](./examples/interface-values/main.go)

# Interface values with nil underlying values

If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.

[See this example.](./examples/interface-values-with-nil/main.go)

# Nil interface values

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

[See this example.](./examples/nil-interface-values/main.go)

# The empty interface

The interface type that specifies zero methods is known as the empty interface:

```
interface{}
```

An empty interface may hold values of any type. (Every type implements at least zero methods.)

**Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.**

[See this example.](./examples/empty-interface/main.go)

# Reference(s)

[A Tour of Go](https://go.dev/tour/methods/9)
