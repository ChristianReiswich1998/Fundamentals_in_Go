package main

import "fmt"

func CountVisits(paths []string) map[string]int {
	// Your code here
	hits := make(map[string]int)

	for _, path := range paths {
		hits[path]++
	}

	return hits
}

func main() {
	log := []string{"/api/users", "/healthz", "/api/users", "/api/orders", "/healthz", "/api/users"}
	fmt.Println(CountVisits(log))
}
