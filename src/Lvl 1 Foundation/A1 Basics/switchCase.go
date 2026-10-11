package main

import "fmt"

func dayType(day string) string {
	// Use a switch on day to return "weekday", "weekend", or "unknown"
	switch day {
	case "Monday":
		return "weekday"
	case "Tuesday":
		return "weekday"
	case "Wednesday":
		return "weekday"
	case "Thursday":
		return "weekday"
	case "Friday":
		return "weekday"
	case "Saturday":
		return "weekend"
	case "Sunday":
		return "weekend"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println(dayType("Monday"))
	fmt.Println(dayType("Saturday"))
}
