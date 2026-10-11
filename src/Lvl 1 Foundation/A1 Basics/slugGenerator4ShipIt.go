package main

import (
	"fmt"
	"strings"
)

// toSlug converts a title into a URL-safe slug
// Use: cleanText, spacesToHyphens, strings.ToLower
func toSlug(title string) string {
	// Your code here
	title = cleanText2(title)
	title = spacesToHyphens(title)
	title = strings.ToLower(title)

	return title
}

func isLetter3(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit3(c rune) bool {
	return c >= '0' && c <= '9'
}

func cleanText2(s string) string {
	result := ""
	for _, ch := range s {
		if isLetter3(ch) || isDigit3(ch) || ch == ' ' {
			result += string(ch)
		}
	}
	return result
}

func spacesToHyphens(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	return strings.Join(strings.Fields(s), "-")
}

func main() {
	fmt.Println(toSlug("My First Blog Post!"))
	fmt.Println(toSlug("Go 1.21: What's New?"))
	fmt.Println(toSlug("  Hello   World  "))
	fmt.Println(toSlug("UPPERCASE THINGS"))
}
