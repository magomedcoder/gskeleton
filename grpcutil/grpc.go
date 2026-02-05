package grpcutil

import (
	"context"
	"errors"
	"fmt"
	"github.com/magomedcoder/gutil/jwtutil"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const GuardGrpcAuth = "grpc-auth"

func GrpcToken(ctx context.Context, storage jwtutil.IStorage, secret string) (*jwtutil.AuthClaims, *string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil, errors.New("не удалось получить метаданные")
	}

	tokens := md.Get("Authorization")
	if len(tokens) == 0 {
		return nil, nil, errors.New("недействительный токен")
	}

	token := strings.TrimPrefix(tokens[len(tokens)-1], "Bearer ")

	userClaims, err := jwtutil.Verify(GuardGrpcAuth, secret, token)
	if err != nil {
		return nil, nil, err
	}

	if storage.IsBlackList(ctx, token) {
		return nil, nil, errors.New("недействительный токен")
	}

	return userClaims, &token, nil
}

func ClientIp(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ipList := md["x-forwarded-for"]; len(ipList) > 0 {
			return ipList[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		addr := p.Addr.String()
		return strings.Split(addr, ":")[0]
	}

	return ""
}

func UserAgent(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("не удалось извлечь данные из контекста")
		return ""
	}

	userAgent := md.Get("user-agent")
	if len(userAgent) == 0 {
		userAgent[0] = "unknown"
	}

	return userAgent[0]
}

func UserId(ctx context.Context) int {
	session, ok := ctx.Value(jwtutil.JWTSession).(*jwtutil.JSession)
	if !ok {
		fmt.Println("Не удалось получить пользователя из контекста")
		return 0
	}

	return session.Uid
}
