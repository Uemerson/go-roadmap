package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "UserID", 123)

	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	userID := ctx.Value("UserID").(int)
	fmt.Println("Processing request for User ID:", userID)
}
