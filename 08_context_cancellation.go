package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Context Cancellation Example ===")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results := make(chan string)

	// Start a goroutine that does work
	go doWork(ctx, results)

	// Wait for result or timeout
	select {
	case result := <-results:
		fmt.Printf("Work completed: %s\n", result)
	case <-ctx.Done():
		fmt.Printf("Work cancelled: %v\n", ctx.Err())
	}

	fmt.Println("Context cancellation example completed")
}

func doWork(ctx context.Context, results chan<- string) {
	select {
	case <-time.After(3 * time.Second): // Simulate long work
		results <- "Work done"
	case <-ctx.Done():
		fmt.Println("Work was cancelled")
		return
	}
}
