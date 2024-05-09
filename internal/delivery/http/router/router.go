package router

import (
	"github.com/gin-gonic/gin"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/http/middleware"
)

func NewRouter(h *handler.Handler, m *middleware.Middleware) *gin.Engine {
	r := gin.New()

	// V1
	newV1(r, h, m)

	return r
}
