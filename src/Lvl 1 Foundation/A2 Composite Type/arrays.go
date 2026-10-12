package main

import "fmt"

func StatusCategory(code int) string {
	// Your code here
	arr := [5]string{"Informational", "Success", "Redirection", "Client Error", "Server Error"}

	if code < 100 || code >= 600 {
		return "Unknown"
	}

	index := (code / 100) - 1

	return arr[index]
}

func main() {
	// Define your categories array here

	fmt.Println(StatusCategory(200))
	fmt.Println(StatusCategory(404))
	fmt.Println(StatusCategory(600))
}
