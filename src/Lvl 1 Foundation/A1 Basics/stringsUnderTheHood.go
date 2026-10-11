package main

import (
	"fmt"
	"unicode/utf8"
)

func firstChar(s string) string {
	// 1. If s is empty, return "" right away.
	// 2. Decode the first rune (import "unicode/utf8").
	// 3. Convert that rune to a string and return it.

	r, size := utf8.DecodeRuneInString(s)

	if size == 0 {
		return ""
	}

	return string(r)

}

func main() {
	fmt.Println(firstChar("Hello"))
	fmt.Println(firstChar("世界"))
	fmt.Println(firstChar("🚀 launch"))
	fmt.Println(firstChar(""))
}
