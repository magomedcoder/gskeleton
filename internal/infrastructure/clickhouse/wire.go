package clickhouse

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/infrastructure/clickhouse/repository"
)

var ProviderSet = wire.NewSet(
	wire.Bind(new(repository.IUserLogRepository), new(*repository.UserLogRepository)),
	repository.NewUserLogRepository,
)
