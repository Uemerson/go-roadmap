package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// properly handling cancellation
				return
			default:
				// do work
			}
		}
	}(ctx)

	time.Sleep(1 * time.Second)

	cancel() // cancel the context
}

func cancel() {
	_, cancel := context.WithCancel(context.Background())
	cancel() // cancel the context
}
