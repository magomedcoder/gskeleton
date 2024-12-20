package ginutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerFunc(fn func(ctx *Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(NewContext(c)); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &Response{
				Message: err.Error(),
			})

			return
		}
	}
}
