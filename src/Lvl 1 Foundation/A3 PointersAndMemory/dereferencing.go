package main

import "fmt"

func main() {
	delay := 200
	p := &delay

	// Read the delay through p (not delay directly), double it, and print the result
	doubleP := 2 * *p
	fmt.Println(doubleP)
}
