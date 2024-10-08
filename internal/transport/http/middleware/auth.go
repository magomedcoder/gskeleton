package middleware

import (
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
		//token := AuthHeaderToken(ctx)
		//if token == "" {
		//	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "Невалидный токен"})
		//	return
		//}
		//
		//fmt.Println(token)

		ctx.Next()
	}
}

func AuthHeaderToken(ctx *gin.Context) string {
	token := ctx.GetHeader("Authorization")
	token = strings.TrimSpace(token)

	return token
}
