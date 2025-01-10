package http

import (
	"github.com/gin-gonic/gin"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v1"
	v2 "github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v2"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/router"
	clickhouseRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/clickhouse/repository"
	postgresRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/repository"
	redisRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/redis/repository"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"github.com/magomedcoder/gskeleton/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupRouter() *gin.Engine {
	postgres := provider.NewPostgresClient(test.GetConfig())
	redis := provider.NewRedisClient(test.GetConfig())
	clickHouse := provider.NewClickHouseClient(test.GetConfig())
	userRepository := postgresRepo.NewUserRepository(postgres)
	userCacheRepository := redisRepo.NewUserCacheRepository(redis)
	userLogRepository := clickhouseRepo.NewUserLogRepository(clickHouse)
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
	r := router.NewRouter(handlerHandler, middlewareMiddleware)
	return r
}

func TestGetUsersHandler(t *testing.T) {
	r := SetupRouter()

	req, _ := http.NewRequest("GET", "/v1/users", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `[{"id":1,"name":"Test"},{"id":2,"name":"Test2"}]`
	assert.JSONEq(t, expected, w.Body.String())
}
