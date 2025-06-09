package app

import (
	"log"
	"os"

	"github.com/magomedcoder/gskeleton/internal/config"
	cliV2 "github.com/urfave/cli/v2"
)

type App struct {
	app *cliV2.App
}

type Action func(ctx *cliV2.Context, conf *config.Config) error

type Command struct {
	Name        string
	Usage       string
	Flags       []cliV2.Flag
	Action      Action
	Subcommands []Command
}

func NewApp(app *cliV2.App) *App {
	return &App{
		app: app,
	}
}

func (a *App) Register(commands ...Command) {
	for _, command := range commands {
		a.app.Commands = append(a.app.Commands, a.createCommand(command))
	}
}

func (a *App) createCommand(cm Command) *cliV2.Command {
	cd := &cliV2.Command{
		Name:        cm.Name,
		Usage:       cm.Usage,
		Flags:       cm.Flags,
		Subcommands: a.createSubcommands(cm.Subcommands),
	}

	if len(cm.Flags) == 0 {
		cd.Flags = a.app.Flags
	} else {
		cd.Flags = cm.Flags
	}

	if cm.Action != nil {
		cd.Action = func(ctx *cliV2.Context) error {
			return cm.Action(ctx, config.New(ctx.String("config")))
		}
	}

	return cd
}

func (a *App) createSubcommands(commands []Command) []*cliV2.Command {
	var subcommands []*cliV2.Command
	for _, subCmd := range commands {
		subcommands = append(subcommands, a.createCommand(subCmd))
	}

	return subcommands
}

func (a *App) Run() {
	if err := a.app.Run(os.Args); err != nil {
		log.Fatalf("Ошибка cli app: %s", err)
	}
}
