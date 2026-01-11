package handler

import (
	"context"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/middleware"
	postgresModel "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type UserOption struct {
	UserUseCase     usecase.IUserUseCase
	TokenMiddleware *middleware.TokenMiddleware
}

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	opts UserOption
}

func NewUserHandler(opts UserOption) *UserHandler {
	return &UserHandler{opts: opts}
}
func (u *UserHandler) Get(ctx context.Context, in *pb.Get_Request) (*pb.Get_Response, error) {
	user, _ := u.opts.UserUseCase.GetUserByUsername(ctx, in.Username)
	if user.Id == 0 {
		return nil, status.Error(codes.NotFound, "Пользователь не найден")
	}

	return &pb.Get_Response{
		Username: user.Username,
		Id:       user.Id,
		CreateAt: user.CreatedAt.Local().String(),
	}, nil
}

func (u *UserHandler) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *UserHandler) Create(ctx context.Context, in *pb.Create_Request) (*pb.Create_Response, error) {
	passwordHash, err := u.opts.UserUseCase.HashPassword(in.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось хешировать пароль")
	}

	user := postgresModel.User{
		Username:  in.Username,
		Password:  passwordHash,
		CreatedAt: time.Now(),
	}

	if _, err = u.opts.UserUseCase.Create(ctx, &user); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Create_Response{
		Success: true,
	}, nil
}
