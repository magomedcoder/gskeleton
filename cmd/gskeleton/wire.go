//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/cli"
	"github.com/magomedcoder/gskeleton/internal/cli/commands"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc"
	"github.com/magomedcoder/gskeleton/internal/delivery/http"
)

func NewHttpInjector(conf *config.Config) *http.AppProvider {
	panic(wire.Build(http.ProviderSet))
}

func NewGrpcInjector(conf *config.Config) *grpc.AppProvider {
	panic(wire.Build(grpc.ProviderSet))
}

func NewCliInjector(conf *config.Config) *commands.AppProvider {
	panic(wire.Build(cli.ProviderSet))
}
