package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/repository"
	userRepo "github.com/magomedcoder/gskeleton/internal/repository/user/repo"
	"github.com/magomedcoder/gskeleton/internal/service"
	userService "github.com/magomedcoder/gskeleton/internal/service/user"
	"github.com/magomedcoder/gskeleton/internal/transport/http/handler"
	"github.com/magomedcoder/gskeleton/internal/transport/http/middleware"
	"github.com/magomedcoder/gskeleton/internal/transport/http/router"
)

type AppProvider struct {
	Engine *gin.Engine
	Config *config.Config
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	provider.NewPostgresDB,

	router.NewRouter,

	handler.ProviderSet,

	wire.Struct(new(middleware.Handler), "*"),
	middleware.NewAuthMiddleware,

	wire.Bind(new(repository.IUserRepository), new(*userRepo.UserRepository)),
	wire.Bind(new(service.IUserService), new(*userService.UserService)),

	userRepo.NewUserRepository,
	userService.NewUserService,
)
