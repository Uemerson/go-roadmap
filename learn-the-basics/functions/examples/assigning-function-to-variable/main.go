package main

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	fmt.Println(area(20, 30))
}
