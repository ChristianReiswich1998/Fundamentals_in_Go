package main

import "fmt"

func Produce(ch chan<- int, n int) {
	// TODO: Send values 1 through n to the channel, then close it
	for i := 1; i <= n; i++ {
		ch <- i
	}

	close(ch)
}

func Consume(ch <-chan int, result chan<- int) {
	// TODO: Sum all received values, send the sum to result channel
	var sum int
	var values []int

	for value := range ch {
		values = append(values, value)
	}

	for i := 0; i < len(values); i++ {
		sum += values[i]
	}

	result <- sum
}

func Transform(in <-chan int, out chan<- int) {
	// TODO: Read each value, multiply by 10, send to output, close output when done
	for i := range in {
		value := i
		value *= 10
		out <- value
	}

	close(out)
}

func main() {
	// Test Produce -> Consume
	ch := make(chan int)
	result := make(chan int)

	go Produce(ch, 5)
	go Consume(ch, result)

	fmt.Println("Sum of 1-5:", <-result) // 15

	// Test Produce -> Transform -> Consume
	ch1 := make(chan int)
	ch2 := make(chan int)
	result2 := make(chan int)

	go Produce(ch1, 5)
	go Transform(ch1, ch2)
	go Consume(ch2, result2)

	fmt.Println("Sum of (1-5)*10:", <-result2) // 150
}
