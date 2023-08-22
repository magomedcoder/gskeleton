//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

type Provider struct{}

func Initialize() *Provider {
	panic(wire.Build(
		wire.NewSet(
			wire.Struct(new(Provider), "*"),
		),
	))
}
