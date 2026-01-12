package handler

import (
	"context"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/domain/entity"
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type UserOption struct {
	UserUseCase usecase.IUserUseCase
}

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	opts UserOption
}

func NewUserHandler(opts UserOption) *UserHandler {
	return &UserHandler{opts: opts}
}

func (u *UserHandler) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &entity.UserOpt{
		Username: in.Username,
		Password: in.Password,
	}

	userModel, err := u.opts.UserUseCase.Create(ctx, user)
	if err != nil {
		log.Printf("Ошибка создания пользователя: %v", err)
		return nil, status.Error(codes.FailedPrecondition, "Ошибка создания пользователя")
	}

	return &pb.CreateUserResponse{
		Id: userModel.Id,
	}, nil
}

func (u *UserHandler) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, _ := u.opts.UserUseCase.GetUserById(ctx, in.Id)
	if user.Id == 0 {
		log.Printf("Пользователь не найден: %v", in.Id)
		return nil, status.Error(codes.NotFound, "Пользователь не найден")
	}

	return &pb.GetUserResponse{
		Id:       user.Id,
		Username: user.Username,
		CreateAt: user.CreatedAt.Unix(),
	}, nil
}
