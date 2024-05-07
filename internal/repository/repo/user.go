package repo

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/repository/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	opts *UserRepositoryOpts
}

type UserRepositoryOpts struct {
	Dao *gorm.DB
}

func NewUserRepository(opts *UserRepositoryOpts) *UserRepository {
	return &UserRepository{opts: opts}
}

func (repo *UserRepository) GetByUsername(username string) (*model.User, error) {
	user := model.User{}

	tx := repo.opts.Dao.Find(&user, "username = ?", username)

	if tx.Error != nil {
		fmt.Printf("Не удалось получить пользователя: %s", tx.Error)
		return nil, tx.Error
	}

	return &user, nil
}

func (repo *UserRepository) CreateUser(user model.User) (*model.User, error) {
	tx := repo.opts.Dao.Create(&user)

	if tx.Error != nil {
		fmt.Printf("Не удалось создать пользователя: %s", tx.Error)
		return &user, tx.Error
	}

	return &user, nil
}
