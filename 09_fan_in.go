package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Fan-in Pattern Example ===")

	// Create multiple producer channels
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	// Start producers
	go producer(ch1, "Producer 1", 3)
	go producer(ch2, "Producer 2", 3)
	go producer(ch3, "Producer 3", 3)

	// Fan-in: merge multiple channels into one
	merged := fanIn(ch1, ch2, ch3)

	// Consume merged messages
	for i := 0; i < 9; i++ {
		fmt.Printf("Consumed: %s\n", <-merged)
	}

	fmt.Println("Fan-in pattern example completed")
}

func producer(ch chan<- string, name string, count int) {
	defer close(ch)
	for i := 1; i <= count; i++ {
		ch <- fmt.Sprintf("%s - Message %d", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func fanIn(inputs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	// Start a goroutine for each input channel
	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for msg := range ch {
				out <- msg
			}
		}(input)
	}

	// Close output channel when all inputs are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
