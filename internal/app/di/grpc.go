package di

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/middleware"
	"github.com/magomedcoder/gskeleton/internal/infrastructure"
	"github.com/magomedcoder/gskeleton/internal/usecase"
)

type GRPCProvider struct {
	Conf            *config.Config
	TokenMiddleware *middleware.TokenMiddleware
	RoutesServices  *middleware.GrpcMethodService
	AuthHandler     pb.AuthServiceServer
	UserHandler     pb.UserServiceServer
}

var GRPCProviderSet = wire.NewSet(
	wire.Struct(new(GRPCProvider), "*"),
	handler.ProviderSet,
	middleware.ProviderSet,
	usecase.ProviderSet,
	infrastructure.ProviderSet,
)
