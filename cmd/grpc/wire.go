//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/repository"
	userRepo "github.com/magomedcoder/gskeleton/internal/repository/user/repo"
	"github.com/magomedcoder/gskeleton/internal/service"
	userService "github.com/magomedcoder/gskeleton/internal/service/user"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/handler"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/router"
)

type Provider struct {
	Server provider.IServer
	//  UserRepository repository.IUserRepository
	//  UserService    service.IUserService
}

var newSet = wire.NewSet(
	wire.Struct(new(Provider), "*"),

	wire.Bind(new(repository.IUserRepository), new(*userRepo.UserRepository)),
	wire.Bind(new(service.IUserService), new(*userService.UserService)),

	//wire.Struct(new(userRepo.UserRepository), "*"),
	//wire.Struct(new(userService.UserService), "*"),
	//wire.Struct(new(middleware.TokenMiddleware), "*"),
	//wire.Struct(new(middleware.AuthMiddleware), "*"),
	//wire.Struct(new(handler.AuthHandler), "*"),
	//wire.Struct(new(handler.UserHandler), "*"),
	//wire.Struct(new(router.GrpcMethodService), "*"),
	//wire.Struct(new(provider.GrpcServer), "*"),

	provider.NewPostgresDB,

	userRepo.NewUserRepository,
	userService.NewUserService,

	middleware.NewTokenMiddleware,
	middleware.NewAuthMiddleware,

	handler.NewAuthHandler,
	handler.NewUserHandler,

	router.NewGrpMethodsService,
	//	middleware.RegisterGlobalService,

	provider.NewGrpcServer,
)

func Initialize(config *config.Config) (Provider, error) {
	wire.Build(newSet)
	return Provider{}, nil
}
