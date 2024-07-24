package repo

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	Dao *gorm.DB
}

func NewUserRepository(
	dao *gorm.DB,
) *UserRepository {
	return &UserRepository{
		Dao: dao,
	}
}
