package repository

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.Bind(new(IUserRepository), new(*UserRepository)),

	NewUserRepository,
)
