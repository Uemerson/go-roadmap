package main

import "fmt"

func main() {
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)
}
