package main

import (
	"github.com/magomedcoder/gskeleton/internal/cli"
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
		Flags: []cliV2.Flag{
			&cliV2.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "/etc/gskeleton/gskeleton.yaml",
				Usage:       "GSkeleton",
				DefaultText: "/etc/gskeleton/gskeleton.yaml",
			},
		},
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return http.Run(ctx, NewHttpInjector(conf))
		},
	}
}

func NewGrpcCommand() provider.Command {
	return provider.Command{
		Name:  "run-grpc",
		Usage: "GRPC server",
		Flags: []cliV2.Flag{
			&cliV2.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "/etc/gskeleton/gskeleton.yaml",
				Usage:       "GSkeleton",
				DefaultText: "/etc/gskeleton/gskeleton.yaml",
			},
		},
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return grpc.Run(ctx, NewGrpcInjector(conf))
		},
	}
}

func NewCliCommand() provider.Command {
	return provider.Command{
		Name:  "cli-migrate",
		Usage: "Cli migrate",
		Flags: []cliV2.Flag{
			&cliV2.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "/etc/gskeleton/gskeleton.yaml",
				Usage:       "GSkeleton",
				DefaultText: "/etc/gskeleton/gskeleton.yaml",
			},
		},
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return cli.Run(ctx, NewCliInjector(conf))
		},
	}
}

func main() {
	app := provider.NewApp(&cliV2.App{
		Name: "GSkeleton",
	})
	app.Register(NewHttpCommand())
	app.Register(NewGrpcCommand())
	app.Register(NewCliCommand())
	app.Run()
}
