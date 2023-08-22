package provider

import "json-rpc-skeleton/pkg/rpc"

func NewRpcServer() *rpc.Server {
	http := &rpc.HTTP{}
	server := rpc.New(
		rpc.WithTransport(http),
	)

	return server
}
