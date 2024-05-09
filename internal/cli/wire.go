package cli

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/provider"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	provider.NewPostgresDB,
)
