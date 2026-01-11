package di

import (
	"github.com/gin-gonic/gin"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v1"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v2"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/router"
)

type HTTPProvider struct {
	Conf   *config.Config
	Engine *gin.Engine
}

func NewHttpInjector(conf *config.Config) *HTTPProvider {
	provider := NewProvider(conf)
	infra := NewInfrastructureProvider(provider)
	useCases := NewUseCaseProvider(infra)

	authMiddleware := middleware.NewAuthMiddleware()
	middlewareMiddleware := &middleware.Middleware{
		AuthMiddleware: authMiddleware,
	}

	user := v1.NewUserHandler(useCases.UserUseCase)
	v2User := v2.NewUserHandler(useCases.UserUseCase)

	handlerHandler := &handler.Handler{
		V1: &v1.V1{
			User: user,
		},
		V2: &v2.V2{
			User: v2User,
		},
	}

	engine := router.NewRouter(handlerHandler, middlewareMiddleware)

	return &HTTPProvider{
		Conf:   conf,
		Engine: engine,
	}
}
