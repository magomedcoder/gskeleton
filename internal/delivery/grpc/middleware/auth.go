package middleware

import (
	"context"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"github.com/magomedcoder/gskeleton/pkg/grpcutil"
	"github.com/magomedcoder/gskeleton/pkg/jwtutil"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthMiddleware struct {
	Conf          *config.Config
	UserUseCase   *usecase.UserUseCase
	PublicMethods map[string]struct{}
}

var publicMethods = map[string]struct{}{
	"/auth.AuthService/Login": {},
}

func NewAuthMiddleware(
	conf *config.Config,
	userUseCase *usecase.UserUseCase,
) *AuthMiddleware {
	return &AuthMiddleware{
		Conf:          conf,
		UserUseCase:   userUseCase,
		PublicMethods: publicMethods,
	}
}

func (a *AuthMiddleware) IsPublicMethod(method string) bool {
	_, exists := a.PublicMethods[method]
	return exists
}

func (a *AuthMiddleware) validateToken(ctx context.Context) (*jwtutil.JSession, error) {
	claims, token, err := grpcutil.GrpcToken(ctx, a.UserUseCase.JwtTokenCacheRepo, a.Conf.Jwt.Secret)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	uid, err := strconv.Atoi(claims.ID)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	return &jwtutil.JSession{
		Uid:       uid,
		Token:     *token,
		ExpiresAt: claims.ExpiresAt.Unix(),
	}, nil
}

func (a *AuthMiddleware) UnaryAuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if a.IsPublicMethod(info.FullMethod) {
		return handler(ctx, req)
	}

	claims, err := a.validateToken(ctx)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, jwtutil.JWTSession, claims)

	return handler(ctx, req)
}

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedStream) Context() context.Context {
	return w.ctx
}

func (a *AuthMiddleware) StreamAuthInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	claims, err := a.validateToken(ss.Context())
	if err != nil {
		return err
	}

	ctx := context.WithValue(ss.Context(), jwtutil.JWTSession, claims)

	wrapped := &wrappedStream{ServerStream: ss, ctx: ctx}
	return handler(srv, wrapped)
}
