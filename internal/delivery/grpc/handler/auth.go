package handler

import (
	"context"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"github.com/magomedcoder/gskeleton/pkg/grpcutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type AuthOption struct {
	UserUseCase usecase.IUserUseCase
}

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer
	opts AuthOption
}

func NewAuthHandler(opts AuthOption) *AuthHandler {
	return &AuthHandler{opts: opts}
}

func (a *AuthHandler) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := a.opts.UserUseCase.Login(ctx, grpcutil.GuardGrpcAuth, in.Username, in.Password)
	if err != nil {
		log.Printf("Ошибка создания токена: %v", err)
		return nil, status.Error(codes.FailedPrecondition, "Ошибка создания токена")
	}

	return &pb.LoginResponse{
		AccessToken: *token,
	}, nil
}
