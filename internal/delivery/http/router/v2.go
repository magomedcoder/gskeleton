package router

import (
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/pkg/core"

	"github.com/gin-gonic/gin"
)

func newV2(r *gin.Engine, h *handler.Handler, m *middleware.Middleware) *gin.Engine {
	authorize := m.AuthMiddleware.Auth()

	v2 := r.Group("/v2")
	{
		user := v2.Group("/users").Use(authorize)
		{
			user.GET("", core.GinHandlerFunc(h.V2.User.List))
			user.GET("/:id", core.GinHandlerFunc(h.V2.User.Get))
		}
	}

	return r
}
