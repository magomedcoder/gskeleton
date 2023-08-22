package rpc

import "context"

type Server struct{}

func New() *Server {
	server := &Server{}

	return server
}

func (r *Server) Run(ctx context.Context) error {
	http := &HTTP{}

	return http.Run()
}
