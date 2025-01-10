package infrastructure

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/infrastructure/clickhouse"
	"github.com/magomedcoder/gskeleton/internal/infrastructure/postgres"
	"github.com/magomedcoder/gskeleton/internal/infrastructure/redis"
)

var ProviderSet = wire.NewSet(
	postgres.ProviderSet,
	redis.ProviderSet,
	clickhouse.ProviderSet,
)
