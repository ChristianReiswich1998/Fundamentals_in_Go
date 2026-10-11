package main

import "fmt"

func main() {
	statusCode := 200

	fmt.Println("Status: " + "OK")
	fmt.Println("Healthy:", statusCode == 200)
	fmt.Println("Latency:", 3.0/2.0)
}
