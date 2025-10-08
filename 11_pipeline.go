package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Pipeline Pattern Example ===")

	// Create pipeline stages
	numbers := generateNumbers(5)
	squares := squareNumbers(numbers)
	results := printResults(squares)

	// Wait for pipeline to complete
	for result := range results {
		fmt.Printf("Final result: %s\n", result)
	}

	fmt.Println("Pipeline pattern example completed")
}

// Stage 1: Generate numbers
func generateNumbers(count int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= count; i++ {
			fmt.Printf("Generating: %d\n", i)
			out <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}

// Stage 2: Square the numbers
func squareNumbers(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			squared := num * num
			fmt.Printf("Squaring %d = %d\n", num, squared)
			out <- squared
		}
	}()
	return out
}

// Stage 3: Format results
func printResults(in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for num := range in {
			result := fmt.Sprintf("Square result: %d", num)
			out <- result
		}
	}()
	return out
}
