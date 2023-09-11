package main

import (
	"context"
	"github.com/urfave/cli/v2"
	"jsonrpc/internal/config"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	cmd := &cli.App{
		Action: func(cCtx *cli.Context) error {
			conf, err := config.ReadConfig(cCtx.Args().Get(0))
			if err != nil {
				return err
			}
			app := Initialize(conf)
			defer cancel()
			if err := app.Server.Run(ctx); err != nil {
				log.Fatal(err)
			}
			return nil
		},
	}
	if err := cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
