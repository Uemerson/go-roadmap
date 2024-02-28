# Marshalling & Unmarshalling JSON

JSON (JavaScript Object Notation) is a simple data interchange format. Syntactically it resembles the objects and lists of JavaScript. It is most commonly used for communication between web back-ends and JavaScript programs running in the browser, but it is used in many other places, too.

# Encoding

To encode JSON data we use the Marshal function.

Given the Go data structure, Message,

```
type Message struct {
    Name string
    Body string
    Time int64
}
```

and an instance of Message

```
m := Message{"Alice", "Hello", 1294706395881547000}
```

we can marshal a JSON-encoded version of m using json.Marshal:

```
b, err := json.Marshal(m)
```

If all is well, err will be nil and b will be a []byte containing this JSON data:

```
b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
```

Only data structures that can be represented as valid JSON will be encoded:

- JSON objects only support strings as keys; to encode a Go map type it must be of the form map[string]T (where T is any Go type supported by the json package).

- Channel, complex, and function types cannot be encoded.

- Cyclic data structures are not supported; they will cause Marshal to go into an infinite loop.

- Pointers will be encoded as the values they point to (or ’null’ if the pointer is nil).

The json package only accesses the exported fields of struct types (those that begin with an uppercase letter). **Therefore only the exported fields of a struct will be present in the JSON output.**

# Reference(s)

[JSON and Go](https://go.dev/blog/json)