package handler

import (
	"context"
	userPb "github.com/magomedcoder/gskeleton/api/grpc/pb/user"
	"github.com/magomedcoder/gskeleton/internal/repository/user/entity"
	"github.com/magomedcoder/gskeleton/internal/service"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type UserHandler struct {
	userPb.UnimplementedUserServiceServer
	UserService     service.IUserService
	TokenMiddleware *middleware.TokenMiddleware
}

func NewUserHandler(
	tokenMiddleware *middleware.TokenMiddleware,
	userService service.IUserService,
) *UserHandler {
	return &UserHandler{
		TokenMiddleware: tokenMiddleware,
		UserService:     userService,
	}
}

func (u *UserHandler) Get(ctx context.Context, in *userPb.Get_Request) (*userPb.Get_Response, error) {
	user, _ := u.UserService.GetUserByUsername(in.Username)
	if user.Id == 0 {
		return nil, status.Error(codes.NotFound, "Пользователь не найден")
	}

	return &userPb.Get_Response{
		Username: user.Username,
		Id:       int32(user.Id),
		CreateAt: user.CreatedAt.Local().String(),
	}, nil
}

func (u *UserHandler) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *UserHandler) Create(ctx context.Context, in *userPb.Create_Request) (*userPb.Create_Response, error) {
	passwordHash, err := u.UserService.HashPassword(in.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось хешировать пароль")
	}

	user := entity.User{
		Username:  in.Username,
		Password:  passwordHash,
		CreatedAt: time.Now(),
	}

	if _, err = u.UserService.Create(user); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &userPb.Create_Response{
		Success: true,
	}, nil
}
