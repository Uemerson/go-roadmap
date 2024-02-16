# Modules

Go modules are a group of related packages that are versioned and distributed together. They specify the requirements of our project, list all the required dependencies, and help us keep track of the specific versions of installed dependencies.

Modules are identified by a module path that is declared in the first line of the go.mod file in our project.

# Packages in Go

A package in Go is the smallest unit of code distribution. They are defined by a directory containing one or more .go source files with the package name declared on top of it.

    All files in the directory must declare the same package name.

Packages allow code to be organized and reused. They provide a way of encapsulating related code into a single unit, which can be imported and used by other packages.

    The Go standard library, for example, consists of many packages such as fmt, os, net, and so on.

**There is a single special package name, main.** This package contains the `main()` function which is the entry point for a project. Every project meant to eventually become an executable must contain the `main()` function, and therefore the main package.

    It is a good practice to declare the main package file(s) at the project's root folder and other packages in their own directories.

# Modules in Go

A module is a collection of related Go packages that are versioned together as a single unit. Modules record precise dependency requirements and create reproducible builds.

A Go module is defined by a go.mod file that resides at the root of the module's directory hierarchy. This file defines the module path, which is the import path prefix for all packages within the module. It specifies the dependencies of the module, including the required versions of other modules.

    To summarize, while a package is a way of structuring and reusing code within a Go program, a module is a versioned collection of packages that also handles dependency management.

# Naming Conventions for Modules

In Go, `module names are used system-wide and therefore should be as specific as possible`, especially if you plan on distributing the module to other developers. The module name is specified in the go.mod file, which acts as the module's manifest and is located at the root of the module's directory hierarchy.

- Module Path: The module path should be a globally unique identifier for the module. It typically takes the form of an internet domain name in reverse order, followed by the module name. For example, github.com/littlejohnny65/example-module. The module path is used as an import path when importing packages from the module.

- Module Name: The module name is the last component of the module path. It should be short, descriptive, and adhere to Go's naming conventions. It is recommended to use lowercase letters with no underscores or mixedCaps. For example, examplemodule

- Versioning: The module name itself does not include version information. The version of a module is specified separately in the go.mod file using a module version identifier, such as v1.2.3. The combination of the module path and the version identifier uniquely identifies a specific version of the module.

It's important to choose meaningful and descriptive names for modules, as they are publicly identifiable and may be used as dependencies in other projects. Clear and consistent naming conventions help in understanding the purpose and context of a module.

# References

[Golang Environment – GOPATH vs go.mod](https://www.freecodecamp.org/news/golang-environment-gopath-vs-go-mod/)
