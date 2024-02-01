package main

import "fmt"

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	fmt.Println(c())
}
