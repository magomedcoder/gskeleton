package repository

import (
	"context"
	"github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/model"
	gormRepo "github.com/magomedcoder/gskeleton/pkg/gorm_repo"
	"gorm.io/gorm"
	"log"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)

	Get(ctx context.Context, id int) (*model.User, error)

	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

var _ IUserRepository = (*UserRepository)(nil)

type UserRepository struct {
	gormRepo.Repo[model.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Repo: gormRepo.NewRepo[model.User](db)}
}

func (u *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if err := u.Repo.Create(ctx, user); err != nil {
		log.Printf("Не удалось создать пользователя: %s", err)
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) Get(ctx context.Context, id int) (*model.User, error) {
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
