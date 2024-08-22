//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/transport/http"
)

func NewHttpInjector(conf *config.Config) *http.AppProvider {
	panic(wire.Build(http.ProviderSet))
}
