package router

import (
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/pkg/ginutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler, m *middleware.Middleware) *gin.Engine {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, ginutil.Response{
			Message: "v1, v2",
		})
	})

	// V1
	newV1(r, h, m)

	// V2
	newV2(r, h, m)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &ginutil.Response{
			Message: "Метод не найден",
		})
	})

	return r
}
