package router

import (
	"json-rpc-skeleton/internal/transport/rpc/handler"
	"json-rpc-skeleton/pkg/rpc"
)

func Methods(server *rpc.Server, handler *handler.Handler) *rpc.Server {

	server.Register("example.get", rpc.Param(handler.Example.Get))

	return server
}
