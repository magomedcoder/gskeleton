package grpc

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/middleware"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/repository"
	"github.com/magomedcoder/gskeleton/internal/usecase"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	provider.NewPostgresDB,
	handler.ProviderSet,
	middleware.ProviderSet,
	usecase.ProviderSet,
	repository.ProviderSet,
)
