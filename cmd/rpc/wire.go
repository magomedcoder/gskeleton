//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"json-rpc-skeleton/internal/config"
	"json-rpc-skeleton/internal/provider"
	"json-rpc-skeleton/internal/transport/rpc/handler"
	"json-rpc-skeleton/pkg/rpc"
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
