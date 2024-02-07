# Type Casting

Go doesn’t support automatic type conversion, but it allows type casting, which is the process of explicitly changing the variable type. To learn more about typecasting

# Type conversions

The expression T(v) converts the value v to the type T.

Some numeric conversions:

```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

Or, put more simply:

```
i := 42
f := float64(i)
u := uint(f)
```

Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions [in the example](./examples/type-conversions/main.go)) and see what happens.

# Reference(s)

[A Tour of Go](https://go.dev/tour/basics/13)
