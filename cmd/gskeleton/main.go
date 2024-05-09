package main

import (
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc"
	"github.com/magomedcoder/gskeleton/internal/delivery/http"
	"github.com/magomedcoder/gskeleton/internal/provider"
	cliV2 "github.com/urfave/cli/v2"
)

func NewHttpCommand() provider.Command {
	return provider.Command{
		Name:  "run-http",
		Usage: "Http server",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return http.Run(ctx, NewHttpInjector(conf))
		},
	}
}

func NewGrpcCommand() provider.Command {
	return provider.Command{
		Name:  "run-grpc",
		Usage: "GRPC server",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return grpc.Run(ctx, NewGrpcInjector(conf))
		},
	}
}

func main() {
	app := provider.NewApp()
	app.Register(NewHttpCommand())
	app.Register(NewGrpcCommand())
	app.Run()
}
