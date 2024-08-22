package router

import (
	"github.com/magomedcoder/gskeleton/internal/transport/http/handler"
	"github.com/magomedcoder/gskeleton/internal/transport/http/middleware"
	"github.com/magomedcoder/gskeleton/pkg/http-server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler, m *middleware.Handler) *gin.Engine {
	r := gin.New()

	authorize := m.AuthMiddleware.Auth()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			post := v1.Group("/example").Use(authorize)
			{
				post.GET("/:id", http_server.HandlerFunc(h.V1.Example.Get))
			}
		}
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &http_server.Response{
			Message: "Метод не найден",
		})
	})

	return r
}
