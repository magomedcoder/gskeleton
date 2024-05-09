package router

import (
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/pkg/http-server"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler, m *middleware.Middleware) *gin.Engine {
	r := gin.New()

	authorize := m.AuthMiddleware.Auth()

	v1 := r.Group("/v1")
	{
		user := v1.Group("/users").Use(authorize)
		{
			user.GET("", http_server.HandlerFunc(h.V1.User.List))
			user.GET("/:id", http_server.HandlerFunc(h.V1.User.Get))
		}
	}

	return r
}
