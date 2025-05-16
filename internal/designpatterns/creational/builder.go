package creational

type Server struct {
	Host string
	Port int
}

type Option func(*Server)

func WithHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

// Builder Pattern: The Builder Pattern creates complex objects through a step-by-step construction
// process and customizable configuration, decoupling the construction process from the object representation
// to achieve a flexible, controllable, and extensible way of object creation.
// You can use Option Builder or Fluent Builder
func NewServer(opts ...Option) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}

	return s
}
