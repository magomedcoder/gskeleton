package repo

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/repository/user/entity"
)

func (repo *UserRepository) Create(user entity.User) (*entity.User, error) {
	tx := repo.Dao.Create(&user)
	if tx.Error != nil {
		fmt.Printf("Не удалось создать пользователя: %s", tx.Error)
		return &user, tx.Error
	}

	return &user, nil
}
