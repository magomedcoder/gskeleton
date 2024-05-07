package middleware

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

type AuthService struct {
	opts *AuthServiceOpts
}

type AuthServiceOpts struct {
	TokenHandler *TokenHandler
}

func NewAuthService(opts *AuthServiceOpts) *AuthService {
	return &AuthService{opts: opts}
}

func (s *AuthService) ValidateToken(ctx context.Context) (*UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("Не удалось получить метаданные")
	}

	token := md.Get("Authorization")
	userClaims, err := s.opts.TokenHandler.ParseToken(token[len(token)-1])
	if err != nil {
		return nil, err
	}
	return userClaims, nil
}
