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

# Composite Types

# Reference(s)

[All data types in Golang with examples](https://golangbyexample.com/all-data-types-in-golang-with-examples/)

[builtin](https://pkg.go.dev/builtin)

[Understanding Data Types in Go](https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go)
