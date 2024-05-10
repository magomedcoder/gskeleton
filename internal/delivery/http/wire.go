package http

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/router"
	"github.com/magomedcoder/gskeleton/internal/infrastructure"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/usecase"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	provider.NewPostgresDB,
	provider.NewRedisClient,
	router.NewRouter,
	handler.ProviderSet,
	middleware.ProviderSet,
	usecase.ProviderSet,
	infrastructure.ProviderSet,
)
