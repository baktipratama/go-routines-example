package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Goroutines WITH WaitGroup ===")

	var wg sync.WaitGroup

	// Launch 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 1; j <= 3; j++ {
				fmt.Printf("Goroutine %d: Message %d\n", id, j)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed")
}
