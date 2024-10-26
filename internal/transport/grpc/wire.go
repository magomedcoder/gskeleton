package grpc

import (
	"github.com/google/wire"
	authPb "github.com/magomedcoder/gskeleton/api/grpc/pb/auth"
	userPb "github.com/magomedcoder/gskeleton/api/grpc/pb/user"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/repository"
	userRepo "github.com/magomedcoder/gskeleton/internal/repository/user/repo"
	"github.com/magomedcoder/gskeleton/internal/service"
	userService "github.com/magomedcoder/gskeleton/internal/service/user"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/handler"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),

	wire.Struct(new(authPb.UnimplementedAuthServiceServer), "*"),
	wire.Struct(new(userPb.UnimplementedUserServiceServer), "*"),

	wire.Struct(new(handler.AuthHandler), "*"),
	wire.Bind(new(authPb.AuthServiceServer), new(*handler.AuthHandler)),

	wire.Struct(new(handler.UserHandler), "*"),
	wire.Bind(new(userPb.UserServiceServer), new(*handler.UserHandler)),

	//	wire.Struct(new(userRepo.UserRepository), "*"),
	wire.Bind(new(repository.IUserRepository), new(*userRepo.UserRepository)),

	//	wire.Struct(new(userService.UserService), "*"),
	wire.Bind(new(service.IUserService), new(*userService.UserService)),

	provider.NewPostgresDB,

	middleware.NewTokenMiddleware,
	middleware.NewGrpMethodsService,

	userService.NewUserService,

	userRepo.NewUserRepository,
)
