package provider

import (
	"json-rpc-skeleton/internal/transport/rpc/handler"
	"json-rpc-skeleton/internal/transport/rpc/router"
	"json-rpc-skeleton/pkg/rpc"
)

func NewRpcServer(handler *handler.Handler) *rpc.Server {
	http := &rpc.HTTP{}
	server := rpc.New(
		rpc.WithTransport(http),
	)
	server = router.Methods(server, handler)

	return server
}
