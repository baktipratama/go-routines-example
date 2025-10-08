package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Goroutines Basic")

	// Launch 3 goroutines
	for i := 1; i <= 3; i++ {
		go func(id int) {
			for j := 1; j <= 3; j++ {
				fmt.Printf("Goroutine %d: Message %d\n", id, j)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Main function ending (goroutines may or may not be done)")
}
