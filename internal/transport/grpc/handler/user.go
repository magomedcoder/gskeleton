package handler

import (
	"context"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/repository/model"
	"github.com/magomedcoder/gskeleton/internal/service"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	opts *UserServerOpts
}

type UserServerOpts struct {
	UserService *service.UserService
	AuthService *middleware.AuthService
}

func NewUserHandler(opts *UserServerOpts) *UserHandler {
	return &UserHandler{opts: opts}
}

func (u *UserHandler) GetUserInfo(ctx context.Context, in *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	user, _ := u.opts.UserService.GetUserByUsername(in.Username)
	if user.Id == 0 {
		return nil, status.Error(codes.NotFound, "Пользователь не найден")
	}

	return &pb.GetUserInfoResponse{
		Username: user.Username,
		Id:       int32(user.Id),
		CreateAt: user.CreatedAt.Local().String(),
	}, nil
}

func (u *UserHandler) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *UserHandler) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	passwordHash, err := u.opts.UserService.HashPassword(in.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Не удалось хешировать пароль")
	}

	user := model.User{
		Username:  in.Username,
		Password:  passwordHash,
		CreatedAt: time.Now(),
	}

	if _, err = u.opts.UserService.CreateUser(user); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateUserResponse{Success: true}, nil
}
