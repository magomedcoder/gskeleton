package router

import (
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/pkg/ginutil"

	"github.com/gin-gonic/gin"
)

func newV1(r *gin.Engine, h *handler.Handler, m *middleware.Middleware) *gin.Engine {
	authorize := m.AuthMiddleware.Auth()

	v1 := r.Group("/v1")
	{
		user := v1.Group("/users").Use(authorize)
		{
			user.POST("", ginutil.HandlerFunc(h.V1.User.CreateUser))
			user.GET("", ginutil.HandlerFunc(h.V1.User.GetUsers))
			user.GET("/:id", ginutil.HandlerFunc(h.V1.User.GetUser))
		}
	}

	return r
}
