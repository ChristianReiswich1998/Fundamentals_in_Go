package main

import (
	"fmt"
	"strconv"
)

type Config struct {
	Port      *int
	EnableTLS *bool
}

type Server struct {
	Port      int
	EnableTLS bool
	Addr      string // listen address, e.g. "localhost:8080"
}

// newServer defaults any unset Config field (Port to 8080, EnableTLS to true),
// then returns a Server holding the resolved values plus its listen Addr
// ("localhost:<port>"). Leave fields the caller already set unchanged.
func newServer(cfg *Config) *Server {
	// Your code here

	if cfg.Port == nil {
		cfg.Port = new(8080)
	}

	if cfg.EnableTLS == nil {
		cfg.EnableTLS = new(true)
	}
	//addr := fmt.Sprintf("localhost:%d", cfg.Port)

	server := &Server{
		Port:      *cfg.Port,
		EnableTLS: *cfg.EnableTLS,
		Addr:      "localhost:" + strconv.Itoa(*cfg.Port),
	}

	return server
}

func (s *Server) start() {
	scheme := "http"
	if s.EnableTLS {
		scheme = "https"
	}
	fmt.Printf("listening on %s://%s\n", scheme, s.Addr)
}

func main() {
	srv := newServer(&Config{})
	srv.start()
}
