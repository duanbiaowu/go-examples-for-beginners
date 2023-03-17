package idiom

import (
	"crypto/tls"
	"time"
)

type Config struct {
	Protocol       string
	Timeout        time.Duration
	MaxConnections int
	TLS            *tls.Config
}

type Server struct {
	Addr string
	Port int
	Conf *Config
}

type Option func(s *Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Conf.Protocol = p
	}
}

func Timeout(t time.Duration) Option {
	return func(s *Server) {
		s.Conf.Timeout = t
	}
}

func MaxConnections(n int) Option {
	return func(s *Server) {
		s.Conf.MaxConnections = n
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.Conf.TLS = tls
	}
}

func NewServer(addr string, port int, options ...Option) (*Server, error) {
	server := &Server{
		Addr: addr,
		Port: port,
		Conf: &Config{},
	}

	for i := range options {
		options[i](server)
	}

	return server, nil
}
