package provider

import "json-rpc-skeleton/pkg/rpc"

func NewRpcServer() *rpc.Server {
	server := rpc.New()

	return server
}
