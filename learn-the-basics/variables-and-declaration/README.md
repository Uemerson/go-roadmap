# Variables in Go

Variable is the name given to a memory location to store a value of a specific type. Go provides multiple ways to declare and use variables.

# Variables

The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level.
[We see both in this example.](./variables/main.go)

# Variables with initializers

A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer. [See this example.](./variables-with-initializers/main.go)

# Short variable declarations

Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

# Basic types

Go's basic types are

    bool

    string

    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr

    byte    // alias for uint8

    rune    // alias for int32
            // represents a Unicode code point

    float32 float64

    complex64 complex128

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

[The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.](./variables-with-initializers/main.go)

# Zero values

Variables declared without an explicit initial value are given their zero value.

The zero value is:

    0 for numeric types,
    false for the boolean type, and
    "" (the empty string) for strings.

[See this example.](./zero-values/main.go)

# Constants

Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

[See this example.](./constants/main.go)

# Numeric Constants

Numeric constants are high-precision values.

An untyped constant takes the type needed by its context.

Try printing needInt(Big) too.

(An int can store at maximum a 64-bit integer, and sometimes less.)

[See this example.](./numeric-constants/main.go)

# Reference(s)

[A Tour of Go](https://go.dev/tour/basics/8)
