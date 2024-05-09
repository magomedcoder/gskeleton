package handler

import (
	"context"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/middleware"
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	TokenMiddleware *middleware.TokenMiddleware
	UserUseCase     usecase.IUserUseCase
}

func (a *AuthHandler) Login(ctx context.Context, in *pb.Login_Request) (*pb.Login_Response, error) {
	user, err := a.UserUseCase.GetUserByUsername(in.Username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Пользователь не найден")
	}

	if _, err := a.UserUseCase.CheckPasswordHash(in.Password, user.Password); err != nil {
		return nil, status.Error(codes.Unauthenticated, "Неверный пароль")
	}

	token, err := a.TokenMiddleware.CreateToken(user)
	if err != nil {
		grpclog.Errorf("Ошибка создания токена: %v", err)
		return nil, status.Error(codes.FailedPrecondition, "ошибка создания токена")
	}

	return &pb.Login_Response{
		AccessToken: token,
	}, nil
}
