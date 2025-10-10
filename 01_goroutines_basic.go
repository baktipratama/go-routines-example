package main

import (
	"fmt"
	"time"
)

// Based on: goroutine.go
func printMessage() {
	fmt.Println("This ran in the background.")
}

func main() {
	go printMessage()
	fmt.Println("This is the main function.")

	// We have to pause here, otherwise the app would close
	// before our helper goroutine gets a chance to run.
	time.Sleep(1 * time.Second)
}
