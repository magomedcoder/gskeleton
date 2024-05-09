package middleware

import (
	"github.com/magomedcoder/gskeleton/pkg/core"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (a *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := AuthHeaderToken(ctx)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &core.Response{
				Message: "Невалидный токен",
			})
			return
		}

		// TODO

		ctx.Next()
	}
}

func AuthHeaderToken(ctx *gin.Context) string {
	token := ctx.GetHeader("Authorization")
	token = strings.TrimSpace(token)

	return token
}
