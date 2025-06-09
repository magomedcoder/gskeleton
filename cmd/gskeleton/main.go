package main

import (
	"github.com/magomedcoder/gskeleton/internal/app"
	"github.com/magomedcoder/gskeleton/internal/app/di"
	"github.com/magomedcoder/gskeleton/internal/cli"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/server"
	cliV2 "github.com/urfave/cli/v2"
)

func main() {
	_app := app.NewApp(&cliV2.App{
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
	_app.Register(
		app.Command{
			Name:  "run",
			Usage: "Run server",

			Subcommands: []app.Command{
				{
					Name:  "http",
					Usage: "Run http server",
					Action: func(ctx *cliV2.Context, conf *config.Config) error {
						return server.HTTP(ctx, di.NewHttpInjector(conf))
					},
				},
				{
					Name:  "grpc",
					Usage: "Run GRPC server",
					Action: func(ctx *cliV2.Context, conf *config.Config) error {
						return server.GRPC(ctx, di.NewGrpcInjector(conf))
					},
				},
			},
		},
		app.Command{
			Name:  "cli",
			Usage: "Cli",
			Subcommands: []app.Command{
				{
					Name:  "migrate",
					Usage: "Migrate",
					Action: func(ctx *cliV2.Context, conf *config.Config) error {
						return cli.RunMigrate(ctx, di.NewCliInjector(conf))
					},
				},
				{
					Name:  "create-user",
					Usage: "Create user",
					Action: func(ctx *cliV2.Context, conf *config.Config) error {
						return cli.RunCreateUser(ctx, di.NewCliInjector(conf))
					},
				},
			},
		},
	)
	_app.Run()
}
