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

# Decoding

To decode JSON data we use the Unmarshal function.

```
func Unmarshal(data []byte, v interface{}) error
```

We must first create a place where the decoded data will be stored

```
var m Message
```

and call json.Unmarshal, passing it a []byte of JSON data and a pointer to m

err := json.Unmarshal(b, &m)

If b contains valid JSON that fits in m, after the call err will be nil and the data from b will have been stored in the struct m, as if by an assignment like:

```
m = Message{
    Name: "Alice",
    Body: "Hello",
    Time: 1294706395881547000,
}
```

How does Unmarshal identify the fields in which to store the decoded data? For a given JSON key "Foo", Unmarshal will look through the destination struct’s fields to find (in order of preference):

- An exported field with a tag of "Foo" (see the [Go spec](https://go.dev/ref/spec#Struct_types) for more on struct tags),

- An exported field named "Foo", or

- An exported field named "FOO" or "FoO" or some other case-insensitive match of "Foo".

What happens when the structure of the JSON data doesn’t exactly match the Go type?

```
b := []byte(`{"Name":"Bob","Food":"Pickle"}`)
var m Message
err := json.Unmarshal(b, &m)
```

Unmarshal will decode only the fields that it can find in the destination type. In this case, only the Name field of m will be populated, and the Food field will be ignored. This behavior is particularly useful when you wish to pick only a few specific fields out of a large JSON blob. It also means that any unexported fields in the destination struct will be unaffected by Unmarshal.

# Nested Objects

Now, consider the case when you have a property called Dimensions, that measures the Height and Length of the bird in question:

```
{
  "species": "pigeon",
  "description": "likes to perch on rocks"
  "dimensions": {
    "height": 24,
    "width": 10
  }
}
```

As with our previous examples, we need to mirror the structure of the object in our Go code. To add a nested `dimensions` object, lets create a `dimensions` struct:

```
type Dimensions struct {
  Height int
  Width int
}
```

Now, the `Bird` struct will include a `Dimensions` field:

```
type Bird struct {
  Species string
  Description string
  Dimensions Dimensions
}
```

We can unmarshal this data using the same method as before:

```
birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
var birds Bird
json.Unmarshal([]byte(birdJson), &birds)
fmt.Printf(bird)
// {pigeon likes to perch on rocks {24 10}}
```

# Primitive Types
We mostly deal with complex objects or arrays when working with JSON, but data like `3`, `3.1412` and `"birds"` are also valid JSON strings.

We can unmarshal these values to their corresponding data type in Go by using primitive types:

```
numberJson := "3"
floatJson := "3.1412"
stringJson := `"bird"`

var n int
var pi float64
var str string

json.Unmarshal([]byte(numberJson), &n)
fmt.Println(n)
// 3

json.Unmarshal([]byte(floatJson), &pi)
fmt.Println(pi)
// 3.1412

json.Unmarshal([]byte(stringJson), &str)
fmt.Println(str)
// bird
```

# Time Values
Did you know that if you try to decode an ISO 8601 date string like `2021-10-18T11:08:47.577Z` into a `time.Time` struct, it will work out of the box?

```
dateJson := `"2021-10-18T11:08:47.577Z"`
var date time.Time
json.Unmarshal([]byte(dateJson), &date)

fmt.Println(date)
// 2021-10-18 11:08:47.577 +0000 UTC
```

Here, dateJson is a JSON string type, but when we unmarshal it into a `time.Time` variable, it is able to understand the JSON data on its own. Well, this is because the `time.Time` struct has a custom UnmarshalJSON method that handles this case.

This will even work if the `time.Time` type is embedded within another struct:

```
type Bird struct {
	Species     string
	Description string
	CreatedAt   time.Time
}

func main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks", "createdAt": "2021-10-18T11:08:47.577Z"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
	// {pigeon likes to perch on rocks 2021-10-18 11:08:47.577 +0000 UTC}
}
```

# Custom Parsing Logic

Similar to the `time.Time` struct, we can also create custom types that implement the `Unmarshaler` interface. This will allow us to define custom logic for decoding JSON data into our custom types.

To illustrate this, let’s take the [nested dimension example](#nested-objects) from before. Suppose we receive the `dimensions` data as a formatted string:

```
{
  "species": "pigeon",
  "description": "likes to perch on rocks",
  "dimensions": "24x10"
}
```

We can modify the `Dimensions` type to implement the Unmarshaler interface, which will have custom parsing logic for our data:

```
type Dimensions struct {
	Height int
	Width  int
}

// unmarshals a JSON string with format
// "heightxwidth" into a Dimensions struct
func (d *Dimensions) UnmarshalJSON(data []byte) error {
	// the "data" parameter is expected to be JSON string as a byte slice
	// for example, `"20x30"`

	if len(data) < 2 {
		return fmt.Errorf("dimensions string too short")
	}
	// remove the quotes
	s := string(data)[1 : len(data)-1]
	// split the string into its two parts
	parts := strings.Split(s, "x")
	if len(parts) != 2 {
		return fmt.Errorf("dimensions string must contain two parts")
	}
	// convert the two parts into ints
	height, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("dimension height must be an int")
	}
	width, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("dimension width must be an int")
	}
	// assign the two ints to the Dimensions struct
	d.Height = height
	d.Width = width
	return nil
}
```

Now, if we try to unmarshal the JSON data, it will create a `Dimensions` struct with the correct values:

```
birdJson := `{"species": "pigeon","description": "likes to perch on rocks", "dimensions":"20x30"}`
var bird Bird
json.Unmarshal([]byte(birdJson), &bird)
fmt.Printf("%+v", bird)
// {Species:pigeon Description:likes to perch on rocks Dimensions:{Height:20 Width:30}}
```

# JSON Struct Tags - Custom Field Names

We saw earlier that Go uses convention to determine the attribute name for mapping JSON properties.

Although sometimes, we want a different attribute name than the one provided in your JSON data. For example, consider the below data:

```
{
  "birdType": "pigeon",
  "what it does": "likes to perch on rocks"
}
```

Here, we would prefer `birdType` to remain as the `Species` attribute in our Go code. It is also not possible for us to provide a suitable attribute name for a key like `"what it does"`.

To solve this, we can use struct field tags:

```
type Bird struct {
  Species string `json:"birdType"`
  Description string `json:"what it does"`
}
```

Now, we can explicitly tell our code which JSON property to map to which attribute.

```
birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
var bird Bird
json.Unmarshal([]byte(birdJson), &bird)
fmt.Println(bird)
// {pigeon likes to perch on rocks}
```

# Reference(s)

[JSON and Go](https://go.dev/blog/json)

[A Complete Guide to JSON in Golang (With Examples)](https://www.sohamkamani.com/golang/json/)