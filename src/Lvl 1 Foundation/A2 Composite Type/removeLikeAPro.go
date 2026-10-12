package main

import (
	"errors"
	"fmt"
	"slices"
)

func RemoveAt(tables []string, index int) ([]string, error) {
	// Your code here
	if index < 0 || index >= len(tables) {
		return nil, errors.New("index out of bounds")
	}

	return slices.Delete(tables, index, index+1), nil
}

func main() {
	tables := []string{"users", "orders", "sessions", "orders", "products"}

	result, err := RemoveAt(tables, 2)
	fmt.Println(result, err)
}
