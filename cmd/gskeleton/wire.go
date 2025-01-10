//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/cli"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc"
	"github.com/magomedcoder/gskeleton/internal/delivery/http"
	"github.com/magomedcoder/gskeleton/internal/provider"
)

var ProviderSet = wire.NewSet(
	provider.NewPostgresClient,
	provider.NewClickHouseClient,
	provider.NewRedisClient,
)

func NewHttpInjector(conf *config.Config) *http.AppProvider {
	panic(wire.Build(ProviderSet, http.ProviderSet))
}

func NewGrpcInjector(conf *config.Config) *grpc.AppProvider {
	panic(wire.Build(ProviderSet, grpc.ProviderSet))
}

func NewCliInjector(conf *config.Config) *cli.AppProvider {
	panic(wire.Build(ProviderSet, cli.ProviderSet))
}
