package main

import "fmt"

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide '%d' by zero", a)
	}
	return a / b, nil
}

func main() {
	value, err := Divide(10, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
