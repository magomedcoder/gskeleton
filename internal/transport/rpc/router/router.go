package router

import (
	"jsonrpc/internal/transport/rpc/handler"
	"jsonrpc/pkg/rpc"
)

func Methods(server *rpc.Server, handler *handler.Handler) *rpc.Server {

	server.Register("example.set", rpc.Param(handler.Example.Set))
	server.Register("example.get", rpc.EmptyParam(handler.Example.Get))

	return server
}
