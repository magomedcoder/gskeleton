package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	app := Initialize()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	if err := app.Server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
