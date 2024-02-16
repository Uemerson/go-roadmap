# Buffer

The `buffer` belongs to the byte package of the Go language, and we can use these package to manipulate the byte of the string.

For example, suppose we have a string. We can read the length of the string with the len function, which will return the numeric length, but what if the strings are too large? We want to calculate in the form of chunks of data, so in such situations, we can use the buffer; the buffer allows us to handle any size of the dynamic length, making it flexible.

# How does a Buffer work in the Go language?

The buffer name itself clarifies its purposes; it allows us to give buffer storage where we can store some data and access the same data when needed. In the case of the string being a large size, there are too many strings that we will store into a variable; then we will create a buffer variable and keep storing the data onto that variable.

We will see the working of the buffer in go language with the help of the below points.

- To use the buffer in the go language, we need to import the bytes package of the go language.
- Once we have imported the bytes package, we can create a variable with the byte package like var x =bytes. Buffer, and on the variable x, we can perform all the operations related to the buffering of string.
- We can store data of string onto the buffer variable x like x.WriteString(“string of message ”), and the data which we stored on the string can be accessed like x.String().
- We can check the length of the string stored on the buffer variable which we have created.
- Even we can store the bytes of data like x.Write([]byte(“Hello “)), and we can get the length of the stored value in the form of a number like x.len().

[See this example](./examples/buffer/main.go)

# Reference(s)

[Golang Buffer](https://www.educba.com/golang-buffer/)
