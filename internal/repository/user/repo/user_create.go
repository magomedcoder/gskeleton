package repo

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/repository/user/model"
)

func (repo *UserRepository) Create(user model.User) (*model.User, error) {
	tx := repo.Dao.Create(&user)
	if tx.Error != nil {
		fmt.Printf("Не удалось создать пользователя: %s", tx.Error)
		return &user, tx.Error
	}

	return &user, nil
}
