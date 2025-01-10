package cli

import (
	"github.com/magomedcoder/gskeleton/internal/cli/handler"
	"github.com/magomedcoder/gskeleton/internal/config"
	cliV2 "github.com/urfave/cli/v2"
)

type AppProvider struct {
	Conf    *config.Config
	Migrate *handler.Migrate
}

func RunMigrate(ctx *cliV2.Context, app *AppProvider) error {
	return app.Migrate.Migrate(ctx.Context)
}

func RunCreateUser(ctx *cliV2.Context, app *AppProvider) error {
	return app.Migrate.CreateUser(ctx.Context)
}
