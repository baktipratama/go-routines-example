package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Buffered Channel Example ===")

	// Based on: buffered-channel.go
	// This channel can hold up to 2 messages before a sender has to wait.
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"

	fmt.Println("All numbers sent to buffered channel")

	// Receive and print the numbers
	fmt.Printf("Received: %s\n", <-messages)
	fmt.Printf("Received: %s\n", <-messages)

	fmt.Println("Buffered channel example completed")
}
