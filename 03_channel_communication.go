package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Channel Communication Example ===")

	// Based on: channel.go
	messages := make(chan string)      // A channel that carries strings.
	go func() { messages <- "ping" }() // Waits until the main function is ready to receive.
	msg := <-messages                  // Waits for the message.
	fmt.Println(msg)

	fmt.Println("Channel communication completed")
}
