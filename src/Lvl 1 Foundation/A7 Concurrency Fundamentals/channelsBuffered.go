package main

import (
	"fmt"
	"sync"
)

// TrySend attempts to send value without blocking.
// Returns true if sent, false if channel is full.
func TrySend(ch chan<- int, value int) bool {
	// TODO: Use select with default for non-blocking send
	// TODO: Attempt to send value to channel without blocking. Return true if sent, false if channel full
	select {
	case ch <- value:
		return true
	default:
		return false
	}
}

// TryReceive attempts to receive without blocking.
// Returns (value, true) if received, (0, false) if channel empty.
func TryReceive(ch <-chan int) (int, bool) {
	// TODO: Use select with default for non-blocking receive
	// TODO: Attempt to receive from channel without blocking. Return (value, true) if received, (0, false) if empty
	var value int

	select {
	case value = <-ch:
		return value, true
	default:
		return 0, false
	}
}

// DrainChannel receives all currently buffered values without blocking.
func DrainChannel(ch <-chan int) []int {
	// TODO: Loop TryReceive until channel is empty
	// TODO: Receive all currently buffered values without blocking. Return as slice
	var result []int
	var receivedValue int
	var channelNotEmpty bool

	for {
		receivedValue, channelNotEmpty = TryReceive(ch)
		if channelNotEmpty {
			result = append(result, receivedValue)
		} else {
			break
		}
	}

	return result
}

// WorkerPool launches numWorkers goroutines to process jobs.
// Each worker reads from jobs, squares the value, sends to results.
// Closes results when all workers are done.
func WorkerPool(jobs <-chan int, results chan<- int, numWorkers int) {
	// TODO: Use WaitGroup to track workers
	// TODO: For each worker: wg.Add(1), launch goroutine with defer wg.Done()
	// TODO: Each worker: for job := range jobs { results <- job*job }
	// TODO: Close results after wg.Wait()
	var waitGroup sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for job := range jobs {
				results <- job * job
			}
		}()
	}
	go func() {
		waitGroup.Wait()
		close(results)
	}()
}

func main() {
	// Test TrySend/TryReceive
	ch := make(chan int, 2)
	fmt.Println("TrySend 1:", TrySend(ch, 1)) // true
	fmt.Println("TrySend 2:", TrySend(ch, 2)) // true
	fmt.Println("TrySend 3:", TrySend(ch, 3)) // false (full)

	v, ok := TryReceive(ch)
	fmt.Printf("TryReceive: %d, %v\n", v, ok) // 1, true

	// Test DrainChannel
	fmt.Println("Drain:", DrainChannel(ch)) // [2]

	// Test WorkerPool
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	go WorkerPool(jobs, results, 2)

	sum := 0
	for r := range results {
		sum += r
	}
	fmt.Println("WorkerPool sum:", sum) // 1+4+9+16+25 = 55
}
