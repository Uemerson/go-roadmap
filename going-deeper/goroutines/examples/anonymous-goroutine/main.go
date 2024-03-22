package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome!! to Main function")

	go func() {
		fmt.Println("Welcome!! to GeeksforGeeks")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("GoodBye!! to Main function")
}
