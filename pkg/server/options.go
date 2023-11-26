package server

type Option func(s *Server)

func WithTransport(transport Transport) Option {
	return func(s *Server) {
		s.transports = append(s.transports, transport)
	}
}
