package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	counter++
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	expectedCounter := 1000

	for i := 0; i < expectedCounter; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Expected Counter:", expectedCounter)
	fmt.Println("Actual Counter:", counter)

	// Check for race condition
	if expectedCounter != counter {
		fmt.Println("Race condition detected!")
	} else {
		fmt.Println("No race condition detected.")
	}
}
