package usecase

import (
	"github.com/magomedcoder/gskeleton/internal/repository/user/model"
	"github.com/magomedcoder/gskeleton/internal/usecase/user"
)

var _ IUserUseCase = (*user.UserUseCase)(nil)

type IUserUseCase interface {
	Create(userModel model.User) (*model.User, error)

	GetUserByUsername(username string) (*model.User, error)

	HashPassword(password string) (string, error)

	CheckPasswordHash(password, hash string) (bool, error)
}
