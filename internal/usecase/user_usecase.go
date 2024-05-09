package usecase

import (
	"context"
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/model"
	"github.com/magomedcoder/gskeleton/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IUserUseCase interface {
	Create(ctx context.Context, userModel *model.User) (*model.User, error)

	GetUserById(ctx context.Context, id int) (*model.User, error)

	GetUserByUsername(ctx context.Context, username string) (*model.User, error)

	HashPassword(password string) (string, error)

	CheckPasswordHash(password, hash string) (bool, error)
}

var _ IUserUseCase = (*UserUseCase)(nil)

type UserUseCase struct {
	UserRepo repository.IUserRepository
}

func (u *UserUseCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *UserUseCase) CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}

func (u *UserUseCase) Create(ctx context.Context, userModel *model.User) (*model.User, error) {
	user, err := u.UserRepo.GetByUsername(ctx, userModel.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось создать пользователя: %s", user.Username))
	}

	if user.Id != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Пользователь %s уже существует", user.Username))
	}

	createdUser, err := u.UserRepo.Create(ctx, userModel)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать пользователя")
	}

	return createdUser, nil
}

func (u *UserUseCase) GetUserById(ctx context.Context, id int) (*model.User, error) {
	user, err := u.UserRepo.Get(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось получить пользователя: %v", id))
	}

	return user, nil
}

func (u *UserUseCase) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	user, err := u.UserRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось получить пользователя: %s", username))
	}

	return user, nil
}
