# Packages

Packages are the most powerful part of the Go language. The purpose of a package is to design and maintain a large number of programs by grouping related features together into single units so that they can be easy to maintain and understand and independent of the other package programs. This modularity allows them to share and reuse. In Go language, every package is defined with a different name and that name is close to their functionality like “strings” package and it contains methods and functions that only related to strings.

# Creating Packages

Packages only really make sense in the context of a separate program which uses them. Without this separate program we have no way of using the package we create.

[In this example.](./examples/packages/main.go):

- math is the name of a package that is part of Go's standard distribution, but since Go packages can be hierarchical we are safe to use the same name for our package. (The real math package is just math, ours is github.com/Uemerson/go-roadmap/learn-the-basics/packages/examples/packages/math)

- When we import our math library we use its full name (import "github.com/Uemerson/go-roadmap/learn-the-basics/packages/examples/packages/math"), but inside of the math.go file we only use the last part of the name (package math).

- We also only use the short name math when we reference functions from our library. If we wanted to use both libraries in the same program Go allows us to use an alias:

```
import m 	"github.com/Uemerson/go-roadmap/learn-the-basics/packages/examples/packages/math"

func main() {
  xs := []float64{1,2,3,4}
  avg := m.Average(xs)
  fmt.Println(avg)
}
```

m is the alias.

- You may have noticed that every function in the packages we've seen start with a capital letter. In Go if something starts with a capital letter that means other packages (and programs) are able to see it. If we had named the function average instead of Average our main program would not have been able to see it.

  It's a good practice to only expose the parts of our package that we want other packages using and hide everything else. This allows us to freely change those parts later without having to worry about breaking other programs, and it makes our package easier to use.

- Package names match the folders they fall in. There are ways around this, but it's a lot easier if you stay within this pattern.

# Explore Go Packages

[Go Packages explorer](https://pkg.go.dev/)

# Giving Names to the Packages

In Go language, when you name a package you must always follow the following points:

- When you create a package the name of the package must be short and simple. For example strings, time, flag, etc. are standard library package.
- The package name should be descriptive and unambiguous.
- Always try to avoid choosing names that are commonly used or used for local relative variables.
- The name of the package generally in the singular form. Sometimes some packages named in plural form like strings, bytes, buffers, etc. Because to avoid conflicts with the keywords.
- Always avoid package names that already have other connotations.

# Reference(s)

[Packages](https://www.golang-book.com/books/intro/11)

[Packages in Golang Read](https://www.geeksforgeeks.org/packages-in-golang/)
