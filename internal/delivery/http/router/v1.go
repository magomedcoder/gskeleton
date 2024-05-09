package router

import (
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/pkg/http-server"

	"github.com/gin-gonic/gin"
)

func newV1(r *gin.Engine, h *handler.Handler, m *middleware.Middleware) *gin.Engine {
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
