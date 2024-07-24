package user

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/repository"
	"github.com/magomedcoder/gskeleton/internal/repository/user/entity"
	def "github.com/magomedcoder/gskeleton/internal/service"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ def.IUserService = (*UserService)(nil)

type UserService struct {
	UserRepo repository.IUserRepository
}

func NewUserService(
	userRepo repository.IUserRepository,
) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *UserService) CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}

func (s *UserService) Create(u entity.User) (*entity.User, error) {
	user, err := s.UserRepo.GetByUsername(u.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось создать пользователя: %s", u.Username))
	}

	if user.Id != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Пользователь %s уже существует", u.Username))
	}

	createdUser, err := s.UserRepo.Create(u)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать пользователя")
	}

	return createdUser, nil
}

func (s *UserService) GetUserByUsername(username string) (*entity.User, error) {
	user, err := s.UserRepo.GetByUsername(username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось получить пользователя: %s", username))
	}

	return user, nil
}
