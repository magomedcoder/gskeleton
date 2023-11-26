package router

import (
	"golang-app-skeleton/internal/transport/app/handler"
	"golang-app-skeleton/pkg/server"
)

func Methods(s *server.Server, h *handler.Handler) *server.Server {

	s.Register("example.set", server.Param(h.Example.Set))
	s.Register("example.get", server.EmptyParam(h.Example.Get))

	return s
}
