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

- **`01_goroutines_basic.go`** - Introduction to launching goroutines with `go` keyword
- **`02_goroutines_with_waitgroup.go`** - Proper synchronization using `sync.WaitGroup` to wait for goroutines
- **`03_channel_communication.go`** - Basic channel operations for sending/receiving data between goroutines
- **`04_buffered_channels.go`** - Channels with internal buffer for asynchronous, non-blocking communication

### Intermediate Examples

- **`05_select_statement.go`** - Non-blocking operations on multiple channels with timeouts
- **`06_worker_pool.go`** - Distributing work among multiple worker goroutines efficiently
- **`07_channel_direction.go`** - Restricting channels to send-only or receive-only for type safety

### Advanced Examples

- **`08_context_cancellation.go`** - Graceful cancellation and timeout handling with context
- **`09_fan_in.go`** - Merging multiple input channels into a single output channel
- **`10_fan_out.go`** - Distributing work from one source to multiple consumers
- **`11_pipeline.go`** - Chaining processing stages to build data transformation pipelines

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