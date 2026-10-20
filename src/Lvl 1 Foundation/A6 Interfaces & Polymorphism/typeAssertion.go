package main

import (
	"errors"
	"fmt"
)

// AsInt safely converts interface{} to int
func AsInt(v interface{}) (int, bool) {
	// TODO: Use type assertion with comma-ok idiom
	// val, ok := v.(int)
	// return val, ok
	val, ok := v.(int)
	if ok == false {
		return 0, false
	}

	return val, ok
}

// AsString safely converts interface{} to string
func AsString(v interface{}) (string, bool) {
	// TODO: Use type assertion with comma-ok idiom
	s, ok := v.(string)
	return s, ok
}

// Add adds two interface{} values if they're both numbers
func Add(a, b interface{}) (float64, error) {
	// TODO: Use type switch to convert a to float64
	// TODO: Use type switch to convert b to float64
	// TODO: Return error if either is not int or float64
	var val_a float64
	var val_b float64

	switch val := a.(type) {
	case int:
		val_a = float64(val)
	case float64:
		val_a = val
	default:
		return 0, errors.New("unsupported types")
	}

	switch val := b.(type) {
	case int:
		val_b = float64(val)
	case float64:
		val_b = val
	default:
		return 0, errors.New("unsupported types")
	}

	return val_a + val_b, nil
}

func main() {
	val, ok := AsInt(42)
	fmt.Printf("AsInt(42): %d, %v\n", val, ok)

	val, ok = AsInt("hello")
	fmt.Printf("AsInt(\"hello\"): %d, %v\n", val, ok)

	str, ok := AsString("world")
	fmt.Printf("AsString(\"world\"): %s, %v\n", str, ok)

	sum, err := Add(10, 20)
	fmt.Printf("Add(10, 20): %.1f, %v\n", sum, err)

	sum, err = Add(10.5, 20.5)
	fmt.Printf("Add(10.5, 20.5): %.1f, %v\n", sum, err)

	sum, err = Add(10, "x")
	fmt.Printf("Add(10, \"x\"): %.1f, %v\n", sum, err)
}
