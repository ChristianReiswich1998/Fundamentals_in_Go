package main

import "fmt"

// cleanText returns s with only letters, digits, and spaces kept
func cleanText(s string) string {
	// Your code here
	result := ""

	for _, ch := range s {
		if isLetter2(ch) || isDigit2(ch) || ch == ' ' {
			result += string(ch)
		}
	}

	return result
}

func isLetter2(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit2(c rune) bool {
	return c >= '0' && c <= '9'
}

func main() {
	fmt.Printf("%q\n", cleanText("Hello, World!"))
	fmt.Printf("%q\n", cleanText("Go 1.21 is great!!!"))
	fmt.Printf("%q\n", cleanText("@#$ test &*"))
	fmt.Printf("%q\n", cleanText(""))
}
