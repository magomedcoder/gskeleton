package cli

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/cli/commands"
	"github.com/magomedcoder/gskeleton/internal/provider"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(commands.AppProvider), "*"),
	provider.NewPostgresDB,
)
