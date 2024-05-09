package handler

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(pb.UnimplementedAuthServiceServer), "*"),
	wire.Struct(new(pb.UnimplementedUserServiceServer), "*"),

	wire.Struct(new(AuthHandler), "*"),
	wire.Bind(new(pb.AuthServiceServer), new(*AuthHandler)),

	wire.Struct(new(UserHandler), "*"),
	wire.Bind(new(pb.UserServiceServer), new(*UserHandler)),
)
