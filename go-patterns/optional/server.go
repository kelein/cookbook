package optional

import (
	"log"
	"time"
)

// Server for requests
type Server struct {
	host    string
	port    int
	maxConn int
	timeout time.Duration
}

// NewServer create a Server instance
func NewServer(opts ...ServerOption) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Start run the server
func (s *Server) Start() error {
	log.Printf("server listen on :%v ...", s.port)
	return nil
}

// ServerOption for Server settings
type ServerOption func(*Server)

// WithHost set server host
func WithHost(host string) ServerOption {
	return func(s *Server) {
		s.host = host
	}
}

// WithPort set server port
func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

// WithMaxConn set server max connection
func WithMaxConn(count int) ServerOption {
	return func(s *Server) {
		s.maxConn = count
	}
}

// WithTimeout set server timeout
func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}
