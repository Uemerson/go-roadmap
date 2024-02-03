# Data Types

Go is a statically typed programming language, which means each variable has a type defined at first and can only hold values with that type. There are two categories of types in Go: basics types and composite types.

## Integers

Integers can be **signed** or **unsigned**.

<details>

<summary open><b>Signed</b> integers are of 5 types as below</summary>

1. int

   **Size:** Platform Dependent.

   - On 32 bit machines, the size of an int will be 32 bits or 4 bytes.
   - On 64 bit machines, the size of an int will be 64 bits or 8 bytes.

   **Range:** Again Platform dependent

   - On 32 bit machines, the range of int will be -2147483648 to 2147483647.
   - On 64 bit machines, the range of int will be -9223372036854775808 to 9223372036854775807.

   **When to Use:**

   - It is a good idea to use int whenever using signed Integer other than the cases mentioned below

     - When the machine is a 32 bit and range needed is greater than -2147483648 to 2147483647, then use int64 instead int. Note that in this case for int64, 2 32-bit memory addresses to form a 64-bit number together.

     - When the range is less then use appropriate integer type.

2. int8

   **Size:** 8 bits or 1 byte.

   **Range:** -128 to 127.

   **When to Use:**

   - Use int8 when there it is known that the int range will be between -128 to 127. For temporary values such as loop invariants, it is still advisable to use int even though it might take more space because it is likely to be promoted to int in some operations or library calls.

   - For array values which lies between -128 to 127, is a good use case for using int8. For eg if you are storing ASCII index for lowercase letters then int8 can be used.

   - It is a good idea to use int8 for data values.

3. int16

   **Size:** 16 bits or 2 byte.

   **Range:** -32768 to 32767.

   **When to Use:**

   - Use int16 when there it is known that the int range will be between -32768 to 32767. For temporary values such as loop invariants, it is still advisable to use int even though it might take more space because it is likely to be promoted to int in some operations or library calls.

   - For array values which lies between -32768 to 32767, is a good use case for using int8. For eg if you are storing ASCII index for lowercase letters than int16 can be used.

4. int32

   **Size:** 32 bits or 4 byte.

   **Range:** -2147483648 to 2147483647.

5. int64

   **Size:** 64 bits or 8 byte.

   **Range:** -9223372036854775808 to 9223372036854775807.

   **When to Use:**

   - **int64** is used when the range is higher. For eg **time.Duration** is of type int64.

</details>

<details open>

<summary><b>UnSigned</b> integers are of 5 types as below</summary>

1. uint

   **Size:** Platform Dependent.

   - On 32 bit machines, the size of an int will be 32 bits or 4 bytes.
   - On 64 bit machines, the size of an int will be 64 bits or 8 bytes.

   **Range:** Again Platform dependent

   - On 32 bit machines, the range of int will be 0 to 4294967295.
   - On 64 bit machines, the range of int will be 0 to 18446744073709551615.

   **When to Use:**

   - It is a good idea to use int whenever using signed Integer other than the cases mentioned below

     - When the machine is a 32 bit and range needed is greater than 0 to 4294967295, then use int64 instead int. Note that in this case for int64, 2 32-bit memory addresses to form a 64-bit number together.

     - When the range is less then use appropriate integer type.

2. uint8

   **Size:** 8 bits or 1 byte.

   **Range:** 0 to 255.

   **When to Use:**

   - Use uint8 when there it is known that the int range will be between 0 to 255. For temporary values such as loop invariants, it is still advisable to use int even though it might take more space because it is likely to be promoted to int in some operations or library calls.

   - For array values which lies between 0 to 255 is a good use case for using uint8. For eg if you are storing ASCII index in an array then uint8 can be used.

3. uint16

   **Size:** 16 bits or 2 byte.

   **Range:** 0 to 65535.

   **When to Use:**

   - Use int16 when there it is known that the int range will be between 0 to 65535. For temporary values such as loop invariants, it is still advisable to use int even though it might take more space because it is likely to be promoted to int in some operations or library calls.

   - For array values which lies between 0 to 65535, is a good use case for using int8.

4. uint32

   **Size:** 32 bits or 4 byte.

   **Range:** 0 to 4294967295.

5. uint64

   **Size:** 64 bits or 8 byte.

   **Range:** 0 to 18446744073709551615.

   **When to Use:**

   - uint64 is used when the range is higher.

**uintptr:**

This is an unsigned integer type that is large enough to hold any pointer address. Therefore is size and range are platform dependent.

**Size:** Platform Dependent.

- On 32 bit machines, the size of int will be 32 bits or 4 byte.
- On 64 bit machines, the size of int will be 64 bits or 8 byte.

**Range:** Again Platform dependent

- On 32 bit machines, the range of int will be 0 to 4294967295.
- On 64 bit machines, the range of int will be 0 to 18446744073709551615.

**Properties**

- A uintptr can be converted to unsafe. Pointer and viceversa.
- Arithmetic can be performed on the uintptr.
- uintptr even though it holds a pointer address, is just a value and does not references any object. Therefore:
  - Its value will not be updated if the corresponding object moves. Eg When goroutine stack changes.
  - The corresponding object can be garbage collected.

**When to Use:**

- Its purpose is to be used along with unsafe.Pointer mainly used for unsafe memory access.

- When you want to save the pointer address value for printing it or storing it. Since the address is just stored and does not reference anything, the corresponding object can be garbage collected.

</details>

<br />

# Float

Floats are numbers with decimals. It is of two types

1. float32

   float32 uses single-precision floating point format to store values. Basically it is the set of all IEEE-754 32-bit floating-point numbers. The 32 bits are divided into – 1 bit sign, 8 bits exponent, and 23 bits mantissa. float 32 take half much size as float 64 and are comparatively faster on some machine architectures.

   **Size:** 32 bits or 4 bytes.

   **Range:** 1.2E-38 to 3.4E+38.

   **DefaultValue:** 0.0

   **When to Use:**

   - If in your system memory is a bottleneck and range is less, then **float32** can be used.

2. float64

   float64 uses a double-precision floating-point format to store values. Basically it is the set of all IEEE-754 64-bit floating-point numbers. The 64 bits are divided into – 1-bit sign, 11 bits exponent, 52 bits mantissa. float64 takes twice as much size compared to float32 but can represent numbers more accurately than float32.

   **Size:** 64 bits or 8 bytes.

   **Range:** -1.7e+308 to +1.7e+308.

   **DefaultValue:** 0.0

   **When to Use:**

   - When the precision needed is high

**float64** is the default float type. When you initialize a variable with a decimal value and don’t specify the float type, the default type inferred will be **float64**.

# Complex Numbers

Complex Numbers are of two types

1. complex64

   Both real and imaginary part are float32

   **Size:** Both real and imaginary part are of same size as float32. It is of size 32 bits or 4 bytes.

   **Range:** Both real and imaginary part range is same as float32 i.e 1.2E-38 to 3.4E+38.

2. complex128

   Both real and imaginary part are float64

   **Size:** Both real and imaginary part are of same size as float64. It is of size 64 bits or 8 bytes.

   **Range:** Both real and imaginary part range is same as float64 i.e -1.7E+308 to +1.7E+308.

# Byte

byte in Go is an alias for uint8 meaning it is an integer value. This integer value is of 8 bits and it represents one byte i.e number between 0-255). A single byte therefore can represent ASCII characters. Golang does not have any data type of ‘char’. Therefore

- byte is used to represent the ASCII character

- rune is used to represent all UNICODE characters which include every character that exists. We will study about rune later in this tutorial.

# Rune

rune in Go is an alias for int32 meaning it is an integer value. This integer value is meant to represent a Unicode Code Point. To understand rune you have to know what Unicode is.

**What is UniCode**

Unicode is a superset of ASCII characters which assigns a unique number to every character that exists. This unique number is called Unicode Code Point.

For eg

- Digit 0 is represented as Unicode Point U+0030 (Decimal Value – 48)
- Small Case b is represented as Unicode Point U+0062 (Decimal Value – 98)
- A pound symbol £ is represented as Unicode Point U+00A3 (Decimal Value – 163)

# String

string is a read only slice of bytes in golang. String can be initialized in two ways

- using double quotes “” eg “this”

string in double quotes honors the escape sequences. For eg if the string contains a \n then while printing there will be a new line

- using back quotes ` eg  \`this`

String in back quotes is just a raw string and it does not honor any kind of escape sequences.

Each character in a string will occupy some bytes depending upon encoding used. For eg in utf-8 encoded string, each character will occupy between 1-4 bytes. You can read about utf-8 in this must read famous blog-The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!). In utf-8 , the characters a or b are encoded using 1 byte while the character pound sign £ is encoded using two bytes . Therefore the string “ab£” will output 4 bytes when you will convert the string to byte array and print it like below

# Booleans

The data type is bool and has two possible values true or false.

Default Value: false

Operations:

    AND – &&
    OR  – ||
    Negation – !

# Composite Types

## Non-Reference Types

### Arrays

Arrays in go are values. They are fixed-length sequences of the same type. Since arrays in Go are values, that is why

- When you assign an array to another variable, it copies the entire array

- When you pass an array as an argument to a function, it makes an entire copy of the array instead of passing just the address

### Structs

In GO struct is named collection of fields. These fields can be of different types. Struct acts as a container of related data of heterogeneous data type.

## Reference Types

### Slices

Slices are dynamically sized, reference into the elements of an array. As mentioned above arrays are of fixed size, so slices give a more flexible interface to arrays. A slice is a reference type as it internally references an array. It is internally represented by three fields

- Address to the underlying array
- Length of the slice
- Capacity of the slice

Slices are typed by the type of the underlying array element, not by its length or capacity. Thus two slices will be of the same type if the type of the underlying array is same irrespective of its length and capacity. The built in append function can be used to add more values to the underlying array. If on using append function the length of the slice increases by current capacity, then a new slice is allocated of double the capacity and elements of the current slice are copied to that.

**The built in function len can you used to get the current length of the sliceInitialising a slice**

### Channels

Channels provide synchronization and communication between goroutines. You can think of it as a pipe through which goroutines can send values and receive values. The operation <- is used to send or receive, with direction of arrow specifying the direction of flow of data

Channel are of two types

- Unbuffered Channel- It doesn't have any capacity to hold and values and thus

  - Send on a channel is block unless there is another goroutine to receive.
  - Receive is block until there is another goroutine on the other side to send.

- Buffered Channel- You can specify the size of buffer here and for them

  - Send on a buffer channel only blocks if the buffer is full
  - Receive is the only block is buffer of the channel is empty

A channel holds data of a particular type at a time. While creating a channel, the type of data l has to be specified while initializing a new channel. Channel can be created using make.

### Maps

maps are golang builtin datatype similar to a hash which map key to a value. maps are referenced data types. When you assign one map to another both refer to the same underlying map.

**Zero Value:** zero value of a map is nil

### Pointers

Pointer is a variable that holds a memory address of another variable. **Zero Value:** The zero value of a pointer is nil.

### Functions

n Go function are values and can be passed around like a value. Basically, function can be used as first-order objects and can be passed around. A function has a name, arguments and returns values. Also, note that there are some important differences between method and function in Go.

### Interface

Interface is a type in Go which is a collection of method signatures. Any type which implements all methods of the interface is of that interface type. Zero value of an interface is nil.

**Interface are implemented implicitly**

There is no explicit declaration that a type implements an interface. In fact, in Go there doesn't exist any "implements" keyword similar to Java. A type implements an interface if it implements all the methods of the interface.

It is correct to define a variable of an interface type and we can assign any concrete type value to this variable if the concrete type implements all the methods of the interface.

**Special case of empty interface**

An empty interface has no methods, hence by default all concrete types implement the empty interface. If you write a function that accepts an empty interface then you can pass any type to that function.

# Reference(s)

[All data types in Golang with examples](https://golangbyexample.com/all-data-types-in-golang-with-examples/)

[builtin](https://pkg.go.dev/builtin)

[Understanding Data Types in Go](https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go)
