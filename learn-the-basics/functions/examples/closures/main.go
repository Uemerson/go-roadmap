package main

import "fmt"

func main() {
	l := 20
	b := 30

	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()
}
