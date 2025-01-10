package grpc

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/middleware"
	"github.com/magomedcoder/gskeleton/internal/infrastructure"
	"github.com/magomedcoder/gskeleton/internal/usecase"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	handler.ProviderSet,
	middleware.ProviderSet,
	usecase.ProviderSet,
	infrastructure.ProviderSet,
)
