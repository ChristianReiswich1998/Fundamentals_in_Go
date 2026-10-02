package main

import "fmt"

func main() {
	var i int = 10
	var f float32 = 3.14
	var tr bool = true
	var fa bool = false
	var s string = "hello"
	var p *int = &i

	fmt.Println("Int i = ", i)
	fmt.Println("Float f = ", f)
	fmt.Println("Boolean true = ", tr)
	fmt.Println("Boolean false = ", fa)
	fmt.Println("String = ", s)
	fmt.Println("Int p = ", p)
}
