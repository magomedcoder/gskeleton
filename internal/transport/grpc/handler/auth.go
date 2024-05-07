package handler

import (
	"context"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/service"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	opts AuthServerOptions
}

type AuthServerOptions struct {
	AuthService  *middleware.AuthService
	UserService  *service.UserService
	TokenHandler *middleware.TokenHandler
}

func NewAuthHandler(opts AuthServerOptions) *AuthHandler {
	return &AuthHandler{opts: opts}
}

func (a *AuthHandler) Login(ctx context.Context, in *pb.Login_Request) (*pb.Login_Response, error) {
	user, err := a.opts.UserService.GetUserByUsername(in.Username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Пользователь не найден")
	}

	if _, err := a.opts.UserService.CheckPasswordHash(in.Password, user.Password); err != nil {
		return nil, status.Error(codes.Unauthenticated, "Неверный пароль")
	}

	token, err := a.opts.TokenHandler.CreateToken(user)
	if err != nil {
		grpclog.Errorf("Ошибка создания токена: %v", err)
		return nil, status.Error(codes.FailedPrecondition, "ошибка создания токена")
	}

	return &pb.Login_Response{AccessToken: token}, nil
}
