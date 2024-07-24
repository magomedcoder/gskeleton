package repo

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/repository/user/entity"
)

func (repo *UserRepository) Get(id int) (*entity.User, error) {
	user := entity.User{}
	tx := repo.Dao.Find(&user, "id = ?", id)
	if tx.Error != nil {
		fmt.Printf("Не удалось получить пользователя: %s", tx.Error)
		return nil, tx.Error
	}

	return &user, nil
}

func (repo *UserRepository) GetByUsername(username string) (*entity.User, error) {
	user := entity.User{}
	tx := repo.Dao.Find(&user, "username = ?", username)
	if tx.Error != nil {
		fmt.Printf("Не удалось получить пользователя: %s", tx.Error)
		return nil, tx.Error
	}

	return &user, nil
}
