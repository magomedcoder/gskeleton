package middleware

import (
	"context"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"reflect"
	"time"
)

type GlobalServiceKey string

var (
	_context context.Context

	AuthServiceKey        = GlobalServiceKey("authService")
	GrpcMethodsServiceKey = GlobalServiceKey("grpcMethodsService")

	globalServicesMap = map[reflect.Type]GlobalServiceKey{
		reflect.TypeOf(&AuthService{}):              AuthServiceKey,
		reflect.TypeOf(&router.GrpcMethodService{}): GrpcMethodsServiceKey,
	}
)

func RegisterGlobalService(ctx context.Context, service interface{}) context.Context {
	serviceType := reflect.TypeOf(service)
	if _, ok := globalServicesMap[serviceType]; !ok {
		log.Fatalf("Неизвестный глобальный сервис: %v", serviceType)
	}

	ctx = context.WithValue(ctx, globalServicesMap[serviceType], service)
	_context = ctx
	return ctx
}

func GetGlobalService(k GlobalServiceKey) interface{} {
	v := _context.Value(k)
	if v == nil {
		log.Fatalf("Значение не найдено: %v", k)
	}
	return v
}

func AuthorizationServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	methodService := GetGlobalService(GrpcMethodsServiceKey).(*router.GrpcMethodService)
	if methodService.IsPublicMethod(info.FullMethod) {
		return handler(ctx, req)
	}

	authService := GetGlobalService(AuthServiceKey).(*AuthService)
	_, err := authService.ValidateToken(ctx)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func LoggingServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)

	grpclog.Infof("Запрос - Метод: %s \t Длительность:%s \t Ошибка:%v \n", info.FullMethod, time.Since(start), err)

	return h, err
}
