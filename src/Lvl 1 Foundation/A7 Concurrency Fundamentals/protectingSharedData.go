package main

import (
	"fmt"
	"sync"
)

// Counter is a thread-safe counter
type Counter struct {
	// TODO: Add mu sync.Mutex
	// TODO: Add value int
	mutex sync.Mutex
	value int
}

// Inc increments the counter by 1
func (c *Counter) Inc() {
	// TODO: Lock, increment, unlock (use defer)
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

// Value returns the current counter value
func (c *Counter) Value() int {
	// TODO: Lock, read, unlock (use defer)
	var value int

	c.mutex.Lock()
	value = c.value
	c.mutex.Unlock()

	return value
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup

	for range 100 {
		wg.Go(func() {
			c.Inc()
		})
	}

	wg.Wait()
	fmt.Printf("Counter value: %d\n", c.Value())
}
