package usecase

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/model"
	"github.com/magomedcoder/gskeleton/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IUserUseCase interface {
	Create(userModel model.User) (*model.User, error)

	GetUserByUsername(username string) (*model.User, error)

	HashPassword(password string) (string, error)

	CheckPasswordHash(password, hash string) (bool, error)
}

var _ IUserUseCase = (*UserUseCase)(nil)

type UserUseCase struct {
	UserRepo repository.IUserRepository
}

func NewUserUseCase(
	userRepo repository.IUserRepository,
) *UserUseCase {
	return &UserUseCase{
		UserRepo: userRepo,
	}
}

func (u *UserUseCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *UserUseCase) CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}

func (u *UserUseCase) Create(userModel model.User) (*model.User, error) {
	user, err := u.UserRepo.GetByUsername(userModel.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось создать пользователя: %s", user.Username))
	}

	if user.Id != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Пользователь %s уже существует", user.Username))
	}

	createdUser, err := u.UserRepo.Create(userModel)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать пользователя")
	}

	return createdUser, nil
}

func (u *UserUseCase) GetUserByUsername(username string) (*model.User, error) {
	user, err := u.UserRepo.GetByUsername(username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось получить пользователя: %s", username))
	}

	return user, nil
}
