package main

import "fmt"

func main() {
	total := 95.50
	people := 4

	// Create a variable "share" that splits total between people
	var share = total / float64(people)
	// Hint: you can't divide a float by an int directly in Go!

	fmt.Printf("Each person pays: $%.2f\n", share)
}
