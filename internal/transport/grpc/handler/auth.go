package handler

import (
	"context"
	"github.com/magomedcoder/gskeleton/internal/service"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
	authPb "github.com/magomedcoder/gskeleton/pkg/pb_generated/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	authPb.UnimplementedAuthServiceServer
	AuthMiddleware  *middleware.AuthMiddleware
	TokenMiddleware *middleware.TokenMiddleware
	UserService     service.IUserService
}

func NewAuthHandler(
	authMiddleware *middleware.AuthMiddleware,
	tokenMiddleware *middleware.TokenMiddleware,
	userService service.IUserService,
) *AuthHandler {
	return &AuthHandler{
		AuthMiddleware:  authMiddleware,
		TokenMiddleware: tokenMiddleware,
		UserService:     userService,
	}
}

func (a *AuthHandler) Login(ctx context.Context, in *authPb.Login_Request) (*authPb.Login_Response, error) {
	user, err := a.UserService.GetUserByUsername(in.Username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Пользователь не найден")
	}

	if _, err := a.UserService.CheckPasswordHash(in.Password, user.Password); err != nil {
		return nil, status.Error(codes.Unauthenticated, "Неверный пароль")
	}

	token, err := a.TokenMiddleware.CreateToken(user)
	if err != nil {
		grpclog.Errorf("Ошибка создания токена: %v", err)
		return nil, status.Error(codes.FailedPrecondition, "ошибка создания токена")
	}

	return &authPb.Login_Response{
		AccessToken: token,
	}, nil
}
