//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/transport/json-rpc/handler"
	"github.com/magomedcoder/gskeleton/pkg/json-rpc-server"
)

type Provider struct {
	Server *json_rpc_server.Server
}

var newSet = wire.NewSet(
	wire.Struct(new(Provider), "*"),
	wire.Struct(new(handler.Handler), "*"),
	provider.NewJsonRpcServer,
	handler.NewExampleHandler,
)

func Initialize(conf *config.Config) *Provider {
	panic(wire.Build(newSet))
}
