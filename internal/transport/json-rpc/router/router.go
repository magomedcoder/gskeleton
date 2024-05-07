package router

import (
	"github.com/magomedcoder/gskeleton/internal/transport/json-rpc/handler"
	"github.com/magomedcoder/gskeleton/pkg/json-rpc-server"
)

func Methods(s *json_rpc_server.Server, h *handler.Handler) *json_rpc_server.Server {

	s.Register("example.set", json_rpc_server.Param(h.Example.Set))
	s.Register("example.get", json_rpc_server.EmptyParam(h.Example.Get))

	return s
}
