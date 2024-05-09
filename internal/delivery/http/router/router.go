package router

import (
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
	"github.com/magomedcoder/gskeleton/pkg/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler, m *middleware.Middleware) *gin.Engine {
	r := gin.New()

	// V1
	newV1(r, h, m)

	// V2
	newV2(r, h, m)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, &core.Response{
			Message: "Метод не найден",
		})
	})

	return r
}
