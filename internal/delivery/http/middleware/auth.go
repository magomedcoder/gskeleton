package middleware

import (
	http_server "github.com/magomedcoder/gskeleton/pkg/http-server"
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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &http_server.Response{
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
