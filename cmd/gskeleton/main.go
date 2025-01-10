package main

import (
	"github.com/magomedcoder/gskeleton/internal/cli"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc"
	"github.com/magomedcoder/gskeleton/internal/delivery/http"
	"github.com/magomedcoder/gskeleton/internal/provider"
	cliV2 "github.com/urfave/cli/v2"
)

func main() {
	app := provider.NewApp(&cliV2.App{
		Name: "GSkeleton",
		Flags: []cliV2.Flag{
			&cliV2.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "/etc/gskeleton/gskeleton.yaml",
				Usage:       "GSkeleton",
				DefaultText: "/etc/gskeleton/gskeleton.yaml",
			},
		},
	})
	app.Register(NewRunCommand())
	app.Register(NewCliCommand())
	app.Run()
}

func NewRunCommand() provider.Command {
	return provider.Command{
		Name:  "run",
		Usage: "Run server",

		Subcommands: []provider.Command{
			{
				Name:  "http",
				Usage: "Run http server",
				Action: func(ctx *cliV2.Context, conf *config.Config) error {
					return http.Run(ctx, NewHttpInjector(conf))
				},
			},
			{
				Name:  "grpc",
				Usage: "Run GRPC server",
				Action: func(ctx *cliV2.Context, conf *config.Config) error {
					return grpc.Run(ctx, NewGrpcInjector(conf))
				},
			},
		},
	}
}

func NewCliCommand() provider.Command {
	return provider.Command{
		Name:  "cli",
		Usage: "Cli",
		Subcommands: []provider.Command{
			{
				Name:  "migrate",
				Usage: "Migrate",
				Action: func(ctx *cliV2.Context, conf *config.Config) error {
					return cli.RunMigrate(ctx, NewCliInjector(conf))
				},
			},
			{
				Name:  "create-user",
				Usage: "Create user",
				Action: func(ctx *cliV2.Context, conf *config.Config) error {
					return cli.RunCreateUser(ctx, NewCliInjector(conf))
				},
			},
		},
	}
}
