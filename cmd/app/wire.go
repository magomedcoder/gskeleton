//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"golang-app-skeleton/internal/config"
	"golang-app-skeleton/internal/provider"
	"golang-app-skeleton/internal/transport/app/handler"
	"golang-app-skeleton/pkg/server"
)

type Provider struct {
	Server *server.Server
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
