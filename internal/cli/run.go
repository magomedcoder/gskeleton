package cli

import (
	"github.com/magomedcoder/gskeleton/internal/app/di"
	cliV2 "github.com/urfave/cli/v2"
)

func RunMigrate(ctx *cliV2.Context, app *di.CLIProvider) error {
	return app.Migrate.Migrate(ctx.Context)
}

func RunCreateUser(ctx *cliV2.Context, app *di.CLIProvider) error {
	return app.Migrate.CreateUser(ctx.Context)
}
