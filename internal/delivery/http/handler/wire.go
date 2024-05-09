package handler

import (
	"github.com/google/wire"
	v1 "github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v1"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(Handler), "*"),

	// V1
	wire.Struct(new(v1.V1), "*"),
	v1.NewUserHandler,
)
