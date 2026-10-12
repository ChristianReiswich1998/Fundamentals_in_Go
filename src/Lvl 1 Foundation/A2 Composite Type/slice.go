package main

import "fmt"

func main() {
	// Create a slice called methods with "GET", "POST", "DELETE"
	methods := []string{"GET", "POST", "DELETE"}

	// Print the slice
	fmt.Println(methods)

	// Print the length
	fmt.Println(len(methods))
}
