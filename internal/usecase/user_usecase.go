package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/domain/entity"
	clickhouseModel "github.com/magomedcoder/gskeleton/internal/infrastructure/clickhouse/model"
	clickhouseRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/clickhouse/repository"
	postgresModel "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/model"
	postgresRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/repository"
	redisModel "github.com/magomedcoder/gskeleton/internal/infrastructure/redis/model"
	redisRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/redis/repository"
	"github.com/magomedcoder/gskeleton/pkg/encrypt"
	"github.com/magomedcoder/gskeleton/pkg/jwtutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"time"
)

type IUserUseCase interface {
	Login(ctx context.Context, guard string, username string, password string) (*string, error)

	Create(ctx context.Context, user *entity.UserOpt) (*entity.User, error)

	GetUsers(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.User, error)

	GetUserById(ctx context.Context, id int64) (*postgresModel.User, error)

	GetUserByUsername(ctx context.Context, username string) (*postgresModel.User, error)
}

var _ IUserUseCase = (*UserUseCase)(nil)

type UserUseCase struct {
	Conf                *config.Config
	UserRepo            postgresRepo.IUserRepository
	UserCacheRepository redisRepo.IUserCacheRepository
	JwtTokenCacheRepo   *redisRepo.JwtTokenCacheRepository
	UserLogRepository   clickhouseRepo.IUserLogRepository
}

func (u *UserUseCase) Login(ctx context.Context, guard string, username string, password string) (*string, error) {
	user, err := u.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Пользователь не найден")
	}

	if _, err := encrypt.CheckPasswordHash(password, user.Password); err != nil {
		return nil, status.Error(codes.Unauthenticated, "Неверный пароль")
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(u.Conf.Jwt.ExpiresTime))

	token := jwtutil.GenerateToken(guard, u.Conf.Jwt.Secret, &jwtutil.Options{
		ExpiresAt: jwtutil.NewNumericDate(expiresAt),
		ID:        username,
		IssuedAt:  jwtutil.NewNumericDate(time.Now()),
	})

	return &token, nil
}

func (u *UserUseCase) Create(ctx context.Context, userEntity *entity.UserOpt) (*entity.User, error) {
	user, err := u.UserRepo.GetByUsername(ctx, userEntity.Username)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println(err)
		}
	}

	if user != nil && user.Id != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("Пользователь %s уже существует", user.Username))
	}

	passwordHash, err := encrypt.HashPassword(userEntity.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось хешировать пароль")
	}

	userModel := &postgresModel.User{
		Username: userEntity.Username,
		Password: passwordHash,
	}

	createdUser, err := u.UserRepo.Create(ctx, userModel)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать пользователя")
	}

	if err := u.UserCacheRepository.Set(ctx, "user_cache", redisModel.UserCache{
		Id:       createdUser.Id,
		Username: createdUser.Username,
	}, int64(time.Hour.Seconds())); err != nil {
		fmt.Printf("не удалось кэшировать пользователя: %v\n", err)
	}

	if err := u.UserLogRepository.Create(ctx, &clickhouseModel.UserLog{
		UserId:    createdUser.Id,
		Log:       "user create",
		CreatedAt: time.Now(),
	}); err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать пользователя")
	}

	return &entity.User{
		Id:       createdUser.Id,
		Username: createdUser.Username,
		Name:     createdUser.Name,
	}, nil
}

func (u *UserUseCase) GetUsers(ctx context.Context, arg ...func(*gorm.DB)) ([]*entity.User, error) {
	users, err := u.UserRepo.GetUsers(ctx, arg...)
	if err != nil {
		return nil, err
	}

	items := make([]*entity.User, 0)
	for _, item := range users {
		items = append(items, &entity.User{
			Id:       item.Id,
			Username: item.Username,
			Name:     item.Name,
		})
	}

	return items, nil
}

func (u *UserUseCase) GetUserById(ctx context.Context, id int64) (*postgresModel.User, error) {
	user, err := u.UserRepo.Get(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось получить пользователя: %v", id))
	}

	if err := u.UserLogRepository.Create(ctx, &clickhouseModel.UserLog{
		UserId:    user.Id,
		Log:       "Get user by id",
		CreatedAt: time.Now(),
	}); err != nil {
		return nil, status.Error(codes.Internal, "Не удалось создать пользователя")
	}

	return user, nil
}

func (u *UserUseCase) GetUserByUsername(ctx context.Context, username string) (*postgresModel.User, error) {
	user, err := u.UserRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Не удалось получить пользователя: %s", username))
	}

	return user, nil
}
