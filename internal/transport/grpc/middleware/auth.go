package middleware

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

type AuthMiddleware struct {
	TokenMiddleware *TokenMiddleware
}

func NewAuthMiddleware(
	tokenMiddleware *TokenMiddleware,
) *AuthMiddleware {
	return &AuthMiddleware{
		TokenMiddleware: tokenMiddleware,
	}
}

func (s *AuthMiddleware) ValidateToken(ctx context.Context) (*UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("Не удалось получить метаданные")
	}

	token := md.Get("Authorization")
	userClaims, err := s.TokenMiddleware.ParseToken(token[len(token)-1])
	if err != nil {
		return nil, err
	}
	return userClaims, nil
}
