package di

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/cli/handler"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/infrastructure"
	"github.com/magomedcoder/gskeleton/internal/usecase"
)

type CLIProvider struct {
	Conf    *config.Config
	Migrate *handler.Migrate
}

var CLIProviderSet = wire.NewSet(
	wire.Struct(new(CLIProvider), "*"),
	wire.Struct(new(handler.Migrate), "*"),
	usecase.ProviderSet,
	infrastructure.ProviderSet,
)
