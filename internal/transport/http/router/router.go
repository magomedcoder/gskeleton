package router

import (
	"github.com/magomedcoder/gskeleton/internal/transport/http/handler"
	"github.com/magomedcoder/gskeleton/internal/transport/http/middleware"
	"github.com/magomedcoder/gskeleton/pkg/http-server"
	"net/http"

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

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &http_server.Response{
			Message: "Метод не найден",
		})
	})

	return r
}
