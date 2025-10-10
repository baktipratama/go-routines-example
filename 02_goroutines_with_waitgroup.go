package main

import (
	"fmt"
	"sync"
	"time"
)

// Based on: waitgroup.go
func workerWg(id int, wg *sync.WaitGroup) {
	// "defer" makes sure wg.Done() is called right before this function exits.
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate some work.
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerWg(i, &wg)
	}
	// This line will wait until every worker has called Done().
	wg.Wait()
	fmt.Println("All workers have finished their jobs.")
}
