package main

import "fmt"

func checkSystem(cpuHigh bool, memoryHigh bool, diskFull bool) string {
	status := "OK"
	// Your code here: reassign `status` to "CRITICAL" or "WARNING" based on
	// the rules in the instructions. Leave it as "OK" otherwise.
	if cpuHigh == true && memoryHigh == true {
		status = "CRITICAL"
	} else if cpuHigh == true || memoryHigh == true || diskFull == true {
		status = "WARNING"
	} else {
		status = "OK"
	}

	// Don't change this line.
	return status
}

func main() {
	cpuHigh := true
	memoryHigh := true
	diskFull := false
	fmt.Println(checkSystem(cpuHigh, memoryHigh, diskFull))

	cpuHigh = true
	memoryHigh = false
	diskFull = false
	fmt.Println(checkSystem(cpuHigh, memoryHigh, diskFull))

	cpuHigh = false
	memoryHigh = false
	diskFull = false
	fmt.Println(checkSystem(cpuHigh, memoryHigh, diskFull))
}
