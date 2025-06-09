package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/router"
	"github.com/magomedcoder/gskeleton/internal/infrastructure"
	"github.com/magomedcoder/gskeleton/internal/usecase"
)

type HTTPProvider struct {
	Conf   *config.Config
	Engine *gin.Engine
}

var HTTPProviderSet = wire.NewSet(
	wire.Struct(new(HTTPProvider), "*"),
	router.NewRouter,
	handler.ProviderSet,
	middleware.ProviderSet,
	usecase.ProviderSet,
	infrastructure.ProviderSet,
)
