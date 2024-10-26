package http

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/repository"
	userRepo "github.com/magomedcoder/gskeleton/internal/repository/user/repo"
	"github.com/magomedcoder/gskeleton/internal/service"
	userService "github.com/magomedcoder/gskeleton/internal/service/user"
	"github.com/magomedcoder/gskeleton/internal/transport/http/handler"
	"github.com/magomedcoder/gskeleton/internal/transport/http/middleware"
	"github.com/magomedcoder/gskeleton/internal/transport/http/router"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),

	wire.Struct(new(middleware.Middleware), "*"),

	//	wire.Struct(new(userRepo.UserRepository), "*"),
	wire.Bind(new(repository.IUserRepository), new(*userRepo.UserRepository)),

	//	wire.Struct(new(userService.UserService), "*"),
	wire.Bind(new(service.IUserService), new(*userService.UserService)),

	provider.NewPostgresDB,

	router.NewRouter,

	handler.ProviderSet,

	middleware.NewAuthMiddleware,

	userService.NewUserService,

	userRepo.NewUserRepository,
)
