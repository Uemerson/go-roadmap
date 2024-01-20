package main

import (
	"errors"
	"fmt"
)

func DoSomething() error {
	return errors.New("something didn't work")
}

func main() {
	err := DoSomething()
	fmt.Println(err)
}
