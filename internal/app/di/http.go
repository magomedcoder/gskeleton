package di

import (
	"github.com/gin-gonic/gin"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v1"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v2"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/router"
	clickhouseRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/clickhouse/repository"
	postgresRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/repository"
	redisRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/redis/repository"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/usecase"
)

type HTTPProvider struct {
	Conf   *config.Config
	Engine *gin.Engine
}

func NewHttpInjector(conf *config.Config) *HTTPProvider {
	db := provider.NewPostgresClient(conf)
	userRepository := postgresRepo.NewUserRepository(db)
	client := provider.NewRedisClient(conf)
	userCacheRepository := redisRepo.NewUserCacheRepository(client)
	conn := provider.NewClickHouseClient(conf)
	userLogRepository := clickhouseRepo.NewUserLogRepository(conn)
	userUseCase := &usecase.UserUseCase{
		UserRepo:            userRepository,
		UserCacheRepository: userCacheRepository,
		UserLogRepository:   userLogRepository,
	}
	user := v1.NewUserHandler(userUseCase)
	v1V1 := &v1.V1{
		User: user,
	}
	v2User := v2.NewUserHandler(userUseCase)
	v2V2 := &v2.V2{
		User: v2User,
	}
	handlerHandler := &handler.Handler{
		V1: v1V1,
		V2: v2V2,
	}
	authMiddleware := middleware.NewAuthMiddleware()
	middlewareMiddleware := &middleware.Middleware{
		AuthMiddleware: authMiddleware,
	}
	engine := router.NewRouter(handlerHandler, middlewareMiddleware)
	httpProvider := &HTTPProvider{
		Conf:   conf,
		Engine: engine,
	}

	return httpProvider
}
