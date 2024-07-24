package service

import "github.com/magomedcoder/gskeleton/internal/repository/user/entity"

type IUserService interface {
	Create(u entity.User) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) (bool, error)
}
