package repository

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user model.User) (*model.User, error)

	Get(id int) (*model.User, error)

	GetByUsername(username string) (*model.User, error)
}

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

func (u *UserRepository) Create(user model.User) (*model.User, error) {
	tx := u.Dao.Create(&user)
	if tx.Error != nil {
		fmt.Printf("Не удалось создать пользователя: %s", tx.Error)
		return &user, tx.Error
	}

	return &user, nil
}

func (u *UserRepository) Get(id int) (*model.User, error) {
	user := model.User{}
	tx := u.Dao.Find(&user, "id = ?", id)
	if tx.Error != nil {
		fmt.Printf("Не удалось получить пользователя: %s", tx.Error)
		return nil, tx.Error
	}

	return &user, nil
}

func (u *UserRepository) GetByUsername(username string) (*model.User, error) {
	user := model.User{}
	tx := u.Dao.Find(&user, "username = ?", username)
	if tx.Error != nil {
		fmt.Printf("Не удалось получить пользователя: %s", tx.Error)
		return nil, tx.Error
	}

	return &user, nil
}
