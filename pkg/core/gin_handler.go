package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinHandlerFunc(fn func(ctx *GinContext) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(NewGinContext(c)); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &Response{
				Message: err.Error(),
			})

			return
		}
	}
}
