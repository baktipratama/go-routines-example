package main

import "fmt"

func main() {
	fmt.Println("=== Go Routines and Channels Examples ===")
	fmt.Println()
	fmt.Println("Available examples:")
	fmt.Println("01. Basic Goroutines          - go run 01_basic_goroutines.go")
	fmt.Println("02. Channel Communication     - go run 02_channel_communication.go")
	fmt.Println("03. Buffered Channels         - go run 03_buffered_channels.go")
	fmt.Println("04. Select Statement          - go run 04_select_statement.go")
	fmt.Println("05. Worker Pool               - go run 05_worker_pool.go")
	fmt.Println("06. Channel Direction         - go run 06_channel_direction.go")
	fmt.Println("07. Context Cancellation      - go run 07_context_cancellation.go")
	fmt.Println("08. Fan-in Pattern            - go run 08_fan_in.go")
	fmt.Println("09. Fan-out Pattern           - go run 09_fan_out.go")
	fmt.Println("10. Pipeline Pattern          - go run 10_pipeline.go")
	fmt.Println()
	fmt.Println("Run any example individually with the command shown above.")
}
