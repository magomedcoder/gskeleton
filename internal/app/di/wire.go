//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/provider"
)

var ProviderSet = wire.NewSet(
	provider.NewPostgresClient,
	provider.NewClickHouseClient,
	provider.NewRedisClient,
)

func NewHttpInjector(conf *config.Config) *HTTPProvider {
	panic(wire.Build(ProviderSet, HTTPProviderSet))
}

func NewGrpcInjector(conf *config.Config) *GRPCProvider {
	panic(wire.Build(ProviderSet, GRPCProviderSet))
}

func NewCliInjector(conf *config.Config) *CLIProvider {
	panic(wire.Build(ProviderSet, CLIProviderSet))
}
