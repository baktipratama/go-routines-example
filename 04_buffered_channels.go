package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Buffered Channel Example ===")

	// Create a buffered channel with capacity 3
	numbers := make(chan int, 3)

	// Send data without blocking (up to buffer size)
	numbers <- 1
	numbers <- 2
	numbers <- 3

	fmt.Println("All numbers sent to buffered channel")

	// Receive and print the numbers
	fmt.Printf("Received: %d\n", <-numbers)
	fmt.Printf("Received: %d\n", <-numbers)
	fmt.Printf("Received: %d\n", <-numbers)

	fmt.Println("Buffered channel example completed")
}
