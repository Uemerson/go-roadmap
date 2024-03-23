package main

import (
	"errors"
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup, err *error) {
	defer wg.Done()

	if id == 2 {
		*err = errors.New("an error occurred") // Simulating work with potential errors
		return
	}

	fmt.Println("Worker", id, "completed")
}

func main() {
	numWorkers := 3
	var wg sync.WaitGroup
	var err error

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, &err)
	}

	wg.Wait() // Wait for all workers to complete

	if err != nil {
		fmt.Println("An error occurred:", err)
	} else {
		fmt.Println("All workers finished")
	}
}
