package main

import "fmt"

// Define your Config struct here
type Config struct {
	Host     string
	Port     int
	MaxConns int
}

// Write your NewConfig function here
func NewConfig(host string, port int) Config {
	config := Config{
		Host:     host,
		Port:     port,
		MaxConns: 100,
	}

	return config
}

func main() {
	c := NewConfig("localhost", 8080)
	fmt.Printf("%s:%d max=%d\n", c.Host, c.Port, c.MaxConns)
}
