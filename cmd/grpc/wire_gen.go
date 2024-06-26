// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/repository/repo"
	"github.com/magomedcoder/gskeleton/internal/service"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/handler"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeUserRepository(dao *gorm.DB) *repo.UserRepository {
	userRepositoryOpts := &repo.UserRepositoryOpts{
		Dao: dao,
	}
	userRepository := repo.NewUserRepository(userRepositoryOpts)
	return userRepository
}

func InitializeUserService(ur *repo.UserRepository) *service.UserService {
	userServiceOpts := &service.UserServiceOpts{
		UserRepo: ur,
	}
	userService := service.NewUserService(userServiceOpts)
	return userService
}

func InitializeUserHandler(us *service.UserService, as *middleware.AuthService) *handler.UserHandler {
	userServerOpts := &handler.UserServerOpts{
		UserService: us,
		AuthService: as,
	}
	userHandler := handler.NewUserHandler(userServerOpts)
	return userHandler
}

func InitializeAuthService(th *middleware.TokenHandler) *middleware.AuthService {
	authServiceOpts := &middleware.AuthServiceOpts{
		TokenHandler: th,
	}
	authService := middleware.NewAuthService(authServiceOpts)
	return authService
}

func InitializeAuthHandler(as *middleware.AuthService, us *service.UserService, th *middleware.TokenHandler) *handler.AuthHandler {
	authServerOptions := handler.AuthServerOptions{
		AuthService:  as,
		UserService:  us,
		TokenHandler: th,
	}
	authHandler := handler.NewAuthHandler(authServerOptions)
	return authHandler
}

func InitializeTokenHandler(secret *config.Jwt) *middleware.TokenHandler {
	tokenHandlerOpts := middleware.TokenHandlerOpts{
		JwtConfig: secret,
	}
	tokenHandler := middleware.NewTokenHandler(tokenHandlerOpts)
	return tokenHandler
}
