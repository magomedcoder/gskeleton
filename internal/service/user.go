package service

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/repository/model"
	"github.com/magomedcoder/gskeleton/internal/repository/repo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	opts *UserServiceOpts
}

type UserServiceOpts struct {
	UserRepo *repo.UserRepository
}

func NewUserService(opts *UserServiceOpts) *UserService {
	return &UserService{opts: opts}
}

func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *UserService) CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}

func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	user, err := s.opts.UserRepo.GetByUsername(username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось получить пользователя: %s", username))
	}

	return user, nil
}

func (s *UserService) CreateUser(u model.User) (*model.User, error) {
	user, err := s.opts.UserRepo.GetByUsername(u.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось создать пользователя: %s", u.Username))
	}

	if user.Id != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Пользователь %s уже существует", u.Username))
	}

	createdUser, err := s.opts.UserRepo.CreateUser(u)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать пользователя")
	}

	return createdUser, nil
}
