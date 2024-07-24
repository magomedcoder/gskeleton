package repository

import "github.com/magomedcoder/gskeleton/internal/repository/user/entity"

type IUserRepository interface {
	Create(user entity.User) (*entity.User, error)
	Get(id int) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
}
