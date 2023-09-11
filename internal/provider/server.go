package provider

import (
	"jsonrpc/internal/config"
	"jsonrpc/internal/transport/rpc/handler"
	"jsonrpc/internal/transport/rpc/router"
	"jsonrpc/pkg/rpc"
)

func NewRpcServer(conf *config.Config, handler *handler.Handler) *rpc.Server {
	http := &rpc.HTTP{
		App: conf.App,
	}
	server := rpc.New(
		rpc.WithTransport(http),
	)
	server = router.Methods(server, handler)

	return server
}
