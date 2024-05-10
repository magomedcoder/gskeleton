package postgres

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/repository"
)

var ProviderSet = wire.NewSet(
	//wire.Struct(new(repository.UserRepository), "*"),
	wire.Bind(new(repository.IUserRepository), new(*repository.UserRepository)),
	repository.NewUserRepository,
)
