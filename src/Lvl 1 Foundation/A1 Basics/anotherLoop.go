package main

import "fmt"

func futureBalance(balance int, years int) int {
	total := balance

	// Use for range years and grow total by 10% each year (total += total / 10).
	for range years {
		total += total / 10
	}

	return total
}

func main() {
	fmt.Println(futureBalance(1000, 0))
	fmt.Println(futureBalance(1000, 1))
	fmt.Println(futureBalance(1000, 5))
	fmt.Println(futureBalance(1000, 10))
}
