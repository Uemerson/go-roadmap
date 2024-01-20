package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

func main() {
	number := 4.0
	result, err := Sqrt(number)

	if err == nil {
		fmt.Printf("Square root of %.1f: %.1f\n", number, result)
	}
}
