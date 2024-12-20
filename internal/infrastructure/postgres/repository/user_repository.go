package repository

import (
	"context"
	"github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/gskeleton/pkg/gormutil"
	"gorm.io/gorm"
	"log"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)

	GetUsers(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.User, error)

	Get(ctx context.Context, id int64) (*model.User, error)

	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

var _ IUserRepository = (*UserRepository)(nil)

type UserRepository struct {
	gormutil.Repo[model.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Repo: gormutil.NewRepo[model.User](db)}
}

func (u *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if err := u.Repo.Create(ctx, user); err != nil {
		log.Printf("Не удалось создать пользователя: %s", err)
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUsers(ctx context.Context, arg ...func(*gorm.DB)) ([]*model.User, error) {
	users, err := u.FindAll(ctx, arg...)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := u.Repo.FindById(ctx, id)
	if err != nil {
		log.Printf("Не удалось получить пользователя: %s", err)
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user, err := u.Repo.FindByWhere(context.TODO(), "username = ?", username)
	if err != nil {
		log.Printf("Не удалось получить пользователя: %s", err)
		return nil, err
	}

	return user, nil
}
