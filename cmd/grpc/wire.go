//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/repository/repo"
	"github.com/magomedcoder/gskeleton/internal/service"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/handler"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
	"gorm.io/gorm"
)

func InitializeUserRepository(dao *gorm.DB) *repo.UserRepository {
	wire.Build(wire.NewSet(wire.Struct(new(repo.UserRepositoryOpts), "*"), repo.NewUserRepository))
	return &repo.UserRepository{}
}

func InitializeUserService(ur *repo.UserRepository) *service.UserService {
	wire.Build(wire.NewSet(wire.Struct(new(service.UserServiceOpts), "*"), service.NewUserService))
	return &service.UserService{}
}

func InitializeUserHandler(us *service.UserService, as *middleware.AuthService) *handler.UserHandler {
	wire.Build(wire.NewSet(wire.Struct(new(handler.UserServerOpts), "*"), handler.NewUserHandler))
	return &handler.UserHandler{}
}

func InitializeAuthService(th *middleware.TokenHandler) *middleware.AuthService {
	wire.Build(wire.NewSet(wire.Struct(new(middleware.AuthServiceOpts), "*"), middleware.NewAuthService))
	return &middleware.AuthService{}
}

func InitializeAuthHandler(as *middleware.AuthService, us *service.UserService, th *middleware.TokenHandler) *handler.AuthHandler {
	wire.Build(wire.NewSet(wire.Struct(new(handler.AuthServerOptions), "*"), handler.NewAuthHandler))
	return &handler.AuthHandler{}
}

func InitializeTokenHandler(secret *config.Jwt) *middleware.TokenHandler {
	wire.Build(wire.NewSet(wire.Struct(new(middleware.TokenHandlerOpts), "*"), middleware.NewTokenHandler))
	return &middleware.TokenHandler{}
}
