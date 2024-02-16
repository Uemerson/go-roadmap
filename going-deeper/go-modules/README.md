# Modules

Go modules are a group of related packages that are versioned and distributed together. They specify the requirements of our project, list all the required dependencies, and help us keep track of the specific versions of installed dependencies.

Modules are identified by a module path that is declared in the first line of the go.mod file in our project.

# Packages in Go

A package in Go is the smallest unit of code distribution. They are defined by a directory containing one or more .go source files with the package name declared on top of it.

    All files in the directory must declare the same package name.

Packages allow code to be organized and reused. They provide a way of encapsulating related code into a single unit, which can be imported and used by other packages.

    The Go standard library, for example, consists of many packages such as fmt, os, net, and so on.

There is a single special package name, main. This package contains the `main()` function which is the entry point for a project. Every project meant to eventually become an executable must contain the `main()` function, and therefore the main package.

    It is a good practice to declare the main package file(s) at the project's root folder and other packages in their own directories.

# References

[Golang Environment – GOPATH vs go.mod](https://www.freecodecamp.org/news/golang-environment-gopath-vs-go-mod/)
