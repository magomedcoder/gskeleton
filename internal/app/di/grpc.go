package di

import (
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/middleware"
)

type GRPCProvider struct {
	Conf        *config.Config
	Middleware  middleware.Middleware
	AuthHandler *handler.AuthHandler
	UserHandler *handler.UserHandler
}

func NewGrpcInjector(conf *config.Config) *GRPCProvider {
	provider := NewProvider(conf)
	infra := NewInfrastructureProvider(provider)
	useCases := NewUseCaseProvider(infra)

	authMiddleware := middleware.NewAuthMiddleware(conf, useCases.UserUseCase)
	loggingMiddleware := middleware.NewLoggingMiddleware()

	middlewareMiddleware := middleware.Middleware{
		Auth:    authMiddleware,
		Logging: loggingMiddleware,
	}

	authOption := handler.AuthOption{
		UserUseCase: useCases.UserUseCase,
	}

	userOption := handler.UserOption{
		UserUseCase: useCases.UserUseCase,
	}

	authHandler := handler.NewAuthHandler(authOption)
	userHandler := handler.NewUserHandler(userOption)

	return &GRPCProvider{
		Conf:        conf,
		Middleware:  middlewareMiddleware,
		AuthHandler: authHandler,
		UserHandler: userHandler,
	}
}
