package main

import "fmt"

// Define your Server struct here
type Server struct {
	Host string
	Port int
	Live bool
}

func Summary(s Server) string {
	// Your code here
	status := ""
	if s.Live == true {
		status = "up"
	} else if s.Live == false {
		status = "down"
	}

	return fmt.Sprintf("%s:%d (%s)", s.Host, s.Port, status)
}

func main() {
	s := Server{Host: "api.example.com", Port: 8080, Live: true}
	fmt.Println(Summary(s))
}
