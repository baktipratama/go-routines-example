package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Channel Direction Example ===")

	messages := make(chan string, 2)

	// Send data using send-only channel
	go sendData(messages)

	// Receive data using receive-only channel
	go receiveData(messages)

	time.Sleep(1 * time.Second)
	fmt.Println("Channel direction example completed")
}

// Send-only channel parameter
func sendData(ch chan<- string) {
	ch <- "Hello"
	ch <- "World"
	close(ch)
}

// Receive-only channel parameter
func receiveData(ch <-chan string) {
	for msg := range ch {
		fmt.Printf("Received: %s\n", msg)
	}
}
