package repository

import (
	"github.com/google/wire"
	userRepo "github.com/magomedcoder/gskeleton/internal/repository/user/repo"
)

var ProviderSet = wire.NewSet(
	//	wire.Struct(new(userRepo.UserRepository), "*"),
	wire.Bind(new(IUserRepository), new(*userRepo.UserRepository)),

	userRepo.NewUserRepository,
)
