# Go Routines and Channels Examples

A comprehensive collection of Go concurrency examples demonstrating goroutines, channels, and common concurrency patterns.

## 📋 Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Examples](#examples)
  - [Basic Examples](#basic-examples)
  - [Advanced Examples](#advanced-examples)
- [Key Concepts](#key-concepts)
- [Running Examples](#running-examples)
- [Learning Path](#learning-path)

## 🎯 Overview

This project provides hands-on examples of Go's concurrency primitives:
- **Goroutines**: Lightweight threads managed by Go runtime
- **Channels**: Communication mechanism between goroutines
- **Select Statements**: Non-blocking channel operations
- **Concurrency Patterns**: Common patterns like worker pools, pipelines, fan-in/fan-out

Each example is self-contained and focuses on a specific concept, making it easy to understand and experiment with Go's concurrency features.

## 📋 Prerequisites

- Go 1.21 or higher
- Basic understanding of Go syntax
- Terminal/command line access

Check your Go version:
```bash
go version
```

## 🚀 Quick Start

1. Clone or download this repository
2. Navigate to the project directory:
   ```bash
   cd go-routines-example
   ```
3. Run the index to see all available examples:
   ```bash
   go run main.go
   ```
4. Run any specific example:
   ```bash
   go run 01_goroutines_basic.go
   ```

## 📚 Examples

### Basic Examples

#### 1. Basic Goroutines (`01_goroutines_basic.go`)
**Concept**: Introduction to goroutines

**Key Learning**:
- How to launch goroutines with `go` keyword

```go
go func(id int) {
    // goroutine work here
}(i)
```

#### 2. Goroutines With WaitGroup (`02_goroutines_with_waitgroup.go`)
**Concept**: Proper goroutine synchronization with `sync.WaitGroup`

**Key Learning**:
- Using `sync.WaitGroup` to wait for goroutines to complete
- Understanding concurrent execution with proper synchronization
- How `defer wg.Done()` ensures cleanup

```go
go func(id int) {
    defer wg.Done()
    // goroutine work here
}(i)
wg.Wait() // Proper synchronization
```

#### 3. Channel Communication (`03_channel_communication.go`)
**Concept**: Basic channel operations for goroutine communication

**Key Learning**:
- Creating unbuffered channels
- Sending and receiving data through channels
- Channel closing and range loops

```go
messages := make(chan string)
messages <- "Hello"  // Send
msg := <-messages    // Receive
```

#### 4. Buffered Channels (`04_buffered_channels.go`)
**Concept**: Channels with internal buffer for asynchronous communication

**Key Learning**:
- Difference between buffered and unbuffered channels
- Non-blocking sends up to buffer capacity
- Buffer size considerations

```go
numbers := make(chan int, 3) // Buffer size 3
```

#### 5. Select Statement (`05_select_statement.go`)
**Concept**: Non-blocking operations on multiple channels

**Key Learning**:
- Handling multiple channels simultaneously
- Timeout patterns with `time.After`
- Non-blocking channel operations

```go
select {
case msg := <-ch1:
    // Handle ch1
case msg := <-ch2:
    // Handle ch2
case <-time.After(timeout):
    // Handle timeout
}
```

#### 6. Worker Pool (`06_worker_pool.go`)
**Concept**: Distributing work among multiple worker goroutines

**Key Learning**:
- Creating a pool of worker goroutines
- Job distribution through channels
- Collecting results from workers

```go
for w := 1; w <= numWorkers; w++ {
    go worker(w, jobs, results)
}
```

### Advanced Examples

#### 7. Channel Direction (`07_channel_direction.go`)
**Concept**: Restricting channel operations with directional channels

**Key Learning**:
- Send-only channels (`chan<- type`)
- Receive-only channels (`<-chan type`)
- Function parameter restrictions for type safety

```go
func sendData(ch chan<- string) { } // Send-only
func receiveData(ch <-chan string) { } // Receive-only
```

#### 8. Context Cancellation (`08_context_cancellation.go`)
**Concept**: Graceful cancellation and timeout handling

**Key Learning**:
- Using `context.Context` for cancellation
- Timeout contexts with `context.WithTimeout`
- Handling cancellation in goroutines

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```

#### 9. Fan-in Pattern (`09_fan_in.go`)
**Concept**: Merging multiple input channels into one output channel

**Key Learning**:
- Combining multiple data sources
- Using `sync.WaitGroup` for coordination
- Dynamic channel merging

```go
func fanIn(inputs ...<-chan string) <-chan string {
    // Merge multiple channels into one
}
```

#### 10. Fan-out Pattern (`10_fan_out.go`)
**Concept**: Distributing work from one source to multiple consumers

**Key Learning**:
- Scaling processing with multiple consumers
- Load distribution patterns
- Concurrent processing of single data stream

```go
go consumer(jobs, results1, "Consumer 1")
go consumer(jobs, results2, "Consumer 2")
```

#### 11. Pipeline Pattern (`11_pipeline.go`)
**Concept**: Chaining processing stages with channels

**Key Learning**:
- Building processing pipelines
- Stage isolation and composition
- Stream processing patterns

```go
numbers := generateNumbers(5)
squares := squareNumbers(numbers)
results := formatResults(squares)
```

## 🔑 Key Concepts

### Goroutines
- Lightweight threads (starts with ~2KB stack)
- Managed by Go runtime scheduler
- Use `go` keyword to launch
- Always use synchronization mechanisms

### Channels
- **Unbuffered**: Synchronous communication (blocking)
- **Buffered**: Asynchronous up to buffer size
- **Directional**: Restrict send/receive operations
- **Closing**: Signal no more values will be sent

### Synchronization
- `sync.WaitGroup`: Wait for multiple goroutines
- `context.Context`: Cancellation and timeouts
- Channel operations: Natural synchronization points

### Common Patterns
- **Worker Pool**: Fixed number of workers processing jobs
- **Fan-in**: Multiple producers → Single consumer
- **Fan-out**: Single producer → Multiple consumers
- **Pipeline**: Sequential processing stages

## 🏃‍♂️ Running Examples

### Run Individual Examples
```bash
# Basic examples
go run 01_goroutines_basic.go
go run 02_goroutines_with_waitgroup.go
go run 03_channel_communication.go
go run 04_buffered_channels.go
go run 05_select_statement.go
go run 06_worker_pool.go

# Advanced examples
go run 07_channel_direction.go
go run 08_context_cancellation.go
go run 09_fan_in.go
go run 10_fan_out.go
go run 11_pipeline.go
```

### View Available Examples
```bash
go run main.go
```

## 📖 Learning Path

### Beginner Path
1. Start with `01_goroutines_basic.go` to see the problem with unreliable synchronization
2. Learn proper synchronization with `02_goroutines_with_waitgroup.go`
3. Learn channel basics with `03_channel_communication.go`
4. Understand buffering with `04_buffered_channels.go`
5. Master channel selection with `05_select_statement.go`

### Intermediate Path
6. Implement worker patterns with `06_worker_pool.go`
7. Learn channel restrictions with `07_channel_direction.go`
8. Handle cancellation with `08_context_cancellation.go`

### Advanced Path
9. Master fan-in pattern with `09_fan_in.go`
10. Implement fan-out with `10_fan_out.go`
11. Build pipelines with `11_pipeline.go`

## 💡 Tips for Learning

1. **Run Examples**: Execute each example and observe the output
2. **Modify Code**: Change parameters like number of workers, buffer sizes
3. **Add Logging**: Insert `fmt.Printf` statements to trace execution
4. **Experiment**: Try breaking the code to understand error conditions
5. **Combine Patterns**: Mix different patterns in your own examples

## 🔧 Troubleshooting

### Common Issues
- **Deadlock**: Usually caused by goroutines waiting for each other
- **Race Conditions**: Use proper synchronization mechanisms
- **Resource Leaks**: Always close channels and handle goroutine cleanup

### Debugging Tips
- Use `go run -race` to detect race conditions
- Add timeouts to prevent hanging
- Use `context.Context` for graceful cancellation

## 📚 Further Reading

- [Go Concurrency Patterns](https://go.dev/talks/2012/concurrency.slide)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go Memory Model](https://go.dev/ref/mem)

---

**Happy Learning! 🚀**

*Each example is designed to be educational and practical. Start with the basics and progressively work through more complex patterns.*