package provider

import (
	"os"

	"github.com/magomedcoder/gskeleton/internal/config"
	cliv2 "github.com/urfave/cli/v2"
)

type App struct {
	app *cliv2.App
}

type Action func(ctx *cliv2.Context, conf *config.Config) error

type Command struct {
	Name        string
	Usage       string
	Flags       []cliv2.Flag
	Action      Action
	Subcommands []Command
}

func NewApp() *App {
	return &App{
		app: &cliv2.App{
			Name:  "GSkeleton",
			Usage: "GSkeleton",
		},
	}
}

func (c *App) Register(cm Command) {
	c.app.Commands = append(c.app.Commands, c.command(cm))
}

func (c *App) command(cm Command) *cliv2.Command {
	cd := &cliv2.Command{
		Name:  cm.Name,
		Usage: cm.Usage,
		Flags: make([]cliv2.Flag, 0),
	}

	if len(cm.Subcommands) > 0 {
		for _, v := range cm.Subcommands {
			cd.Subcommands = append(cd.Subcommands, c.command(v))
		}
	} else {
		if cm.Flags != nil && len(cm.Flags) > 0 {
			cd.Flags = append(cd.Flags, cm.Flags...)
		}

		var isConfig bool

		for _, flag := range cd.Flags {
			if flag.Names()[0] == "config" {
				isConfig = true
				break
			}
		}

		if !isConfig {
			cd.Flags = append(cd.Flags, &cliv2.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "/etc/gskeleton/gskeleton.yaml",
				Usage:       "GSkeleton",
				DefaultText: "/etc/gskeleton/gskeleton.yaml",
			})
		}

		if cm.Action != nil {
			cd.Action = func(ctx *cliv2.Context) error {
				return cm.Action(ctx, config.New(ctx.String("config")))
			}
		}
	}

	return cd
}

func (c *App) Run() {
	_ = c.app.Run(os.Args)
}
