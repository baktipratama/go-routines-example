package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Channel Communication Example ===")

	// Create an unbuffered channel
	messages := make(chan string)

	// Send data in a goroutine
	go func() {
		messages <- "Hello"
		messages <- "from"
		messages <- "goroutine"
		close(messages)
	}()

	// Receive data from the channel
	for msg := range messages {
		fmt.Printf("Received: %s\n", msg)
	}

	fmt.Println("Channel communication completed")
}
