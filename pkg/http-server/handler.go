package http_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerFunc(fn func(ctx *Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(New(c)); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &Response{
				Message: err.Error(),
			})

			return
		}
	}
}
