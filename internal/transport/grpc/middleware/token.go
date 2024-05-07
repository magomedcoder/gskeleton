package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/repository/model"
	"time"
)

type TokenHandlerOpts struct {
	JwtConfig *config.Jwt
}

type TokenHandler struct {
	opts TokenHandlerOpts
}

func NewTokenHandler(opts TokenHandlerOpts) *TokenHandler {
	return &TokenHandler{opts: opts}
}

type UserInfo struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
}

type UserClaims struct {
	*jwt.StandardClaims
	UserInfo
}

func (t *TokenHandler) CreateToken(user *model.User) (string, error) {
	claims := &UserClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "github.com/magomedcoder/gskeleton-grpc",
		},
		UserInfo{user.Username, user.Id},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.opts.JwtConfig.Secret))
}

func (t *TokenHandler) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Неожиданный метод подписи: %v", token.Header["alg"])
		}

		return []byte(t.opts.JwtConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
