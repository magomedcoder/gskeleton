package repository

import "github.com/magomedcoder/gskeleton/internal/repository/user/model"

type IUserRepository interface {
	Create(user model.User) (*model.User, error)

	Get(id int) (*model.User, error)

	GetByUsername(username string) (*model.User, error)
}
