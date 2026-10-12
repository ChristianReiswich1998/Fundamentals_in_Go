package main

import (
	"fmt"
	"strings"
)

func FormatLogs(logs []string) string {
	// Your code here
	var stringArr []string

	if len(logs) == 0 {
		return ""
	}

	for i := range logs {
		s := fmt.Sprintf("%d. %s", i+1, logs[i])
		stringArr = append(stringArr, s)
	}

	return strings.Join(stringArr, "\n")
}

func main() {
	result := FormatLogs([]string{"server started", "request received", "response sent"})
	// %q shows the string as a quoted literal: each newline appears
	// as a visible \n so you can spot a stray trailing one
	fmt.Printf("%q\n", result)
	fmt.Println("---")
	// Println shows the same string with real line breaks
	fmt.Println(result)
}
