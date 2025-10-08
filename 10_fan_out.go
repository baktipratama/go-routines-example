package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Fan-out Pattern Example ===")

	// Single producer channel
	jobs := make(chan int, 10)

	// Create multiple consumer channels
	results1 := make(chan string, 5)
	results2 := make(chan string, 5)

	// Start consumers
	go consumer(jobs, results1, "Consumer 1")
	go consumer(jobs, results2, "Consumer 2")

	// Produce jobs
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	// Collect results from both consumers
	time.Sleep(500 * time.Millisecond)
	close(results1)
	close(results2)

	fmt.Println("Results from Consumer 1:")
	for result := range results1 {
		fmt.Printf("  %s\n", result)
	}

	fmt.Println("Results from Consumer 2:")
	for result := range results2 {
		fmt.Printf("  %s\n", result)
	}

	fmt.Println("Fan-out pattern example completed")
}

func consumer(jobs <-chan int, results chan<- string, name string) {
	for job := range jobs {
		result := fmt.Sprintf("%s processed job %d", name, job)
		results <- result
		time.Sleep(50 * time.Millisecond)
	}
}
