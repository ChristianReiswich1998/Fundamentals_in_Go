package main

import "fmt"

func AddLogs(logBuffer []string, newLines []string) (oldCap, newCap int, resized bool) {
	// Your code here
	oldCap = cap(logBuffer)

	if len(newLines) == 0 {
		newCap = oldCap
		resized = false
		return oldCap, newCap, resized
	}

	for i := range newLines {
		logBuffer = append(logBuffer, newLines[i])
	}

	newCap = cap(logBuffer)

	if oldCap == newCap {
		resized = false
	} else if oldCap < newCap {
		resized = true
	}

	return oldCap, newCap, resized
}

func main() {
	spacious := make([]string, 0, 10)
	fmt.Println(AddLogs(spacious, []string{"boot", "ready", "listening"})) // 10 10 false

	tight := make([]string, 0, 2)
	fmt.Println(AddLogs(tight, []string{"boot", "ready", "listening"})) // 2 <bigger> true
}
