package http

import (
	"github.com/gin-gonic/gin"
	"github.com/magomedcoder/gskeleton/internal/app/di"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler/v1"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/router"
	"github.com/magomedcoder/gskeleton/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupRouter() *gin.Engine {
	provider := di.NewProvider(test.GetConfig())
	infra := di.NewInfrastructureProvider(provider)
	useCases := di.NewUseCaseProvider(infra)
	user := v1.NewUserHandler(useCases.UserUseCase)

	v1V1 := &v1.V1{
		User: user,
	}

	handlerHandler := &handler.Handler{
		V1: v1V1,
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
