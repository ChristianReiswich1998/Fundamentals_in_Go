package main

import "fmt"

// Write your function here
func AddPermission(permissions []string, permission string) []string {
	return append(permissions, permission)

}

func main() {
	result := AddPermission([]string{"read"}, "write")
	fmt.Println(result)
}
