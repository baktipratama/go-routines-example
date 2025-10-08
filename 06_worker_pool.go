package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Worker Pool Pattern Example ===")

	const numWorkers = 3
	const numJobs = 9

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= numJobs; r++ {
		result := <-results
		fmt.Printf("Job result: %d\n", result)
	}

	fmt.Println("Worker pool example completed")
}

// Worker function for the worker pool
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- job * 2                 // Send result
	}
}
