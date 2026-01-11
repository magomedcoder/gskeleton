package di

import (
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/middleware"
)

type GRPCProvider struct {
	Conf            *config.Config
	TokenMiddleware *middleware.TokenMiddleware
	RoutesServices  *middleware.GrpcMethodService
	AuthHandler     *handler.AuthHandler
	UserHandler     *handler.UserHandler
}

func NewGrpcInjector(conf *config.Config) *GRPCProvider {
	provider := NewProvider(conf)
	infra := NewInfrastructureProvider(provider)
	useCases := NewUseCaseProvider(infra)

	tokenMiddleware := middleware.NewTokenMiddleware(conf)
	grpcMethodService := middleware.NewGrpMethodsService()

	authOption := handler.AuthOption{
		TokenMiddleware: tokenMiddleware,
		UserUseCase:     useCases.UserUseCase,
	}

	userOption := handler.UserOption{
		UserUseCase:     useCases.UserUseCase,
		TokenMiddleware: tokenMiddleware,
	}

	authHandler := handler.NewAuthHandler(authOption)
	userHandler := handler.NewUserHandler(userOption)

	return &GRPCProvider{
		Conf:            conf,
		TokenMiddleware: tokenMiddleware,
		RoutesServices:  grpcMethodService,
		AuthHandler:     authHandler,
		UserHandler:     userHandler,
	}
}
