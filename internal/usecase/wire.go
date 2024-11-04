package usecase

import (
	"github.com/google/wire"
	userUseCase "github.com/magomedcoder/gskeleton/internal/usecase/user"
)

var ProviderSet = wire.NewSet(
	//	wire.Struct(new(userUseCase.UserUseCase), "*"),
	wire.Bind(new(IUserUseCase), new(*userUseCase.UserUseCase)),

	userUseCase.NewUserUseCase,
)
