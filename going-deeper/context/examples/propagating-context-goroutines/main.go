package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "UserID", 123)

	go performTask(ctx)

	// Continue with other operations
}

func performTask(ctx context.Context) {
	userID := ctx.Value("UserID")
	fmt.Println("User ID:", userID)
}
