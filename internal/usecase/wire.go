package usecase

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	wire.Struct(new(UserUseCase), "*"),
	wire.Bind(new(IUserUseCase), new(*UserUseCase)),
)
