package main

import (
	"fmt"
	"time"
)

func WithTimeout(ch <-chan int, timeout time.Duration) (int, bool) {
	// TODO: Use select with ch and time.After
	// TODO: Return (value, true) if received, (0, false) on timeout
	select {
	case val := <-ch:
		return val, true
	case <-time.After(timeout):
		return 0, false
	}
}

func TryReceive3(ch <-chan int) (int, bool) {
	// TODO: Use select with default
	select {
	case val := <-ch:
		return val, true
	default:
		return 0, false
	}
}

func FanIn(ch1, ch2 <-chan int, done <-chan bool) <-chan int {
	// TODO: Multiplex ch1 and ch2, stop on done
	channel := make(chan int)

	go func() {
		for {
			select {
			case val1 := <-ch1:

				channel <- val1
			case val2 := <-ch2:

				channel <- val2
			case _ = <-done:
				return
			}
		}
	}()

	return channel
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 42
	val, ok := WithTimeout(ch1, 1*time.Second)
	fmt.Printf("WithTimeout (success): %d, %v\n", val, ok)

	ch2 := make(chan int)
	val, ok = WithTimeout(ch2, 100*time.Millisecond)
	fmt.Printf("WithTimeout (timeout): %d, %v\n", val, ok)

	ch3 := make(chan int, 1)
	ch3 <- 99
	val, ok = TryReceive3(ch3)
	fmt.Printf("TryReceive (available): %d, %v\n", val, ok)

	ch4 := make(chan int)
	val, ok = TryReceive3(ch4)
	fmt.Printf("TryReceive (empty): %d, %v\n", val, ok)
}
