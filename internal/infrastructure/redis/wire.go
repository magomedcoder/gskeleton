package redis

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/infrastructure/redis/repository"
)

var ProviderSet = wire.NewSet(
	wire.Bind(new(repository.IUserCacheRepository), new(*repository.UserCacheRepository)),
	repository.NewUserCacheRepository,
)
