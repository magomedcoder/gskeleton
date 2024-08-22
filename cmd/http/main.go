package main

import (
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/transport/http"
	cliv2 "github.com/urfave/cli/v2"
)

func NewHttpCommand() provider.Command {
	return provider.Command{
		Name:  "run",
		Usage: "Http server",
		Action: func(ctx *cliv2.Context, conf *config.Config) error {
			return http.Run(ctx, NewHttpInjector(conf))
		},
	}
}

func main() {
	app := provider.NewApp()
	app.Register(NewHttpCommand())
	app.Run()
}
