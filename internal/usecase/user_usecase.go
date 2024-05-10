package usecase

import (
	"context"
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/domain/entity"
	postgresModel "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/model"
	postgresRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/repository"
	redisModel "github.com/magomedcoder/gskeleton/internal/infrastructure/redis/model"
	redisRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/redis/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

type IUserUseCase interface {
	Create(ctx context.Context, userModel *postgresModel.User) (*postgresModel.User, error)

	GetUsers(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.User, error)

	GetUserById(ctx context.Context, id int64) (*postgresModel.User, error)

	GetUserByUsername(ctx context.Context, username string) (*postgresModel.User, error)

	HashPassword(password string) (string, error)

	CheckPasswordHash(password, hash string) (bool, error)
}

var _ IUserUseCase = (*UserUseCase)(nil)

type UserUseCase struct {
	PostgresUserRepo         postgresRepo.IUserRepository
	RedisUserCacheRepository redisRepo.IUserCacheRepository
}

func (u *UserUseCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *UserUseCase) CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}

func (u *UserUseCase) GetUsers(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.User, error) {
	users, err := u.PostgresUserRepo.GetUsers(ctx, arg...)
	if err != nil {
		return nil, err
	}

	items := make([]*entity.User, 0)
	for _, item := range users {
		items = append(items, &entity.User{
			Id:   item.Id,
			Name: item.Username,
		})
	}

	return items, nil
}

func (u *UserUseCase) Create(ctx context.Context, userModel *postgresModel.User) (*postgresModel.User, error) {
	user, err := u.PostgresUserRepo.GetByUsername(ctx, userModel.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось создать пользователя: %s", user.Username))
	}

	if user.Id != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Пользователь %s уже существует", user.Username))
	}

	createdUser, err := u.PostgresUserRepo.Create(ctx, userModel)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать пользователя")
	}

	if err := u.RedisUserCacheRepository.Set(ctx, "user_cache", redisModel.UserCache{
		Id:       createdUser.Id,
		Username: createdUser.Username,
	}, int64(time.Hour.Seconds())); err != nil {
		fmt.Printf("не удалось кэшировать пользователя: %v\n", err)
	}

	return createdUser, nil
}

func (u *UserUseCase) GetUserById(ctx context.Context, id int64) (*postgresModel.User, error) {
	user, err := u.PostgresUserRepo.Get(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось получить пользователя: %v", id))
	}

	return user, nil
}

func (u *UserUseCase) GetUserByUsername(ctx context.Context, username string) (*postgresModel.User, error) {
	user, err := u.PostgresUserRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось получить пользователя: %s", username))
	}

	return user, nil
}
