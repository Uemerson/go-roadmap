package main

import (
	"fmt"

	math "github.com/Uemerson/go-roadmap/learn-the-basics/packages/examples/packages/mypackage"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
