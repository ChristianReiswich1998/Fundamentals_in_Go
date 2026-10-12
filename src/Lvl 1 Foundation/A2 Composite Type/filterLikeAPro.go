package main

import (
	"fmt"
	"slices"
	"strings"
)

func DropTempTables(tables []string) []string {
	// Your code here
	isHidden := func(name string) bool {
		return strings.HasPrefix(name, "temp_")
	}

	return slices.DeleteFunc(tables, isHidden)
}

func main() {
	tables := []string{"users", "temp_001", "orders", "temp_cache", "products"}
	cleaned := DropTempTables(tables)
	fmt.Println(cleaned)
}
