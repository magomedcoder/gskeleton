//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"jsonrpc/internal/config"
	"jsonrpc/internal/provider"
	"jsonrpc/internal/transport/rpc/handler"
	"jsonrpc/pkg/rpc"
)

type Provider struct {
	Server *rpc.Server
}

var newSet = wire.NewSet(
	wire.Struct(new(Provider), "*"),
	wire.Struct(new(handler.Handler), "*"),
	provider.NewRpcServer,
	handler.NewExampleHandler,
)

func Initialize(conf *config.Config) *Provider {
	panic(wire.Build(newSet))
}
