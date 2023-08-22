//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"json-rpc-skeleton/internal/provider"
)

type Provider struct{}

var newSet = wire.NewSet(
	wire.Struct(new(Provider), "*"),
	provider.NewRpcServer,
)

func Initialize() *Provider {
	panic(wire.Build(newSet))
}
