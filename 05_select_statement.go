package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Select Statement Example ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()

	// Use select to handle multiple channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("Timeout occurred")
		}
	}

	fmt.Println("Select statement example completed")
}
