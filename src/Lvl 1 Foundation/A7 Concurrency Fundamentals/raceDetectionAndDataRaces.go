package main

import (
	"sync"
)

// Counter is a thread-safe counter
type Counter3 struct {
	// TODO: Add a sync.Mutex field
	mutex sync.Mutex
	value int
}

// Inc increments the counter by 1
func (c *Counter) Inc3() {
	// TODO: Lock the mutex before incrementing
	// TODO: Use defer to unlock
	c.mutex.Lock()
	c.value++
	defer c.mutex.Unlock()
}

// Value returns the current counter value
func (c *Counter) Value3() int {
	// TODO: Lock the mutex before reading
	// TODO: Use defer to unlock
	var value int

	c.mutex.Lock()
	value = c.value
	defer c.mutex.Unlock()

	return value
}

func main() {

}
