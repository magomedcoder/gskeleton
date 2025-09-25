package di

import (
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/handler"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/middleware"
	clickhouseRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/clickhouse/repository"
	postgresRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/repository"
	redisRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/redis/repository"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/usecase"
)

type GRPCProvider struct {
	Conf            *config.Config
	TokenMiddleware *middleware.TokenMiddleware
	RoutesServices  *middleware.GrpcMethodService
	AuthHandler     pb.AuthServiceServer
	UserHandler     pb.UserServiceServer
}

func NewGrpcInjector(conf *config.Config) *GRPCProvider {
	tokenMiddleware := middleware.NewTokenMiddleware(conf)
	grpcMethodService := middleware.NewGrpMethodsService()
	unimplementedAuthServiceServer := pb.UnimplementedAuthServiceServer{}
	db := provider.NewPostgresClient(conf)
	userRepository := postgresRepo.NewUserRepository(db)
	client := provider.NewRedisClient(conf)
	userCacheRepository := redisRepo.NewUserCacheRepository(client)
	conn := provider.NewClickHouseClient(conf)
	userLogRepository := clickhouseRepo.NewUserLogRepository(conn)
	userUseCase := &usecase.UserUseCase{
		UserRepo:            userRepository,
		UserCacheRepository: userCacheRepository,
		UserLogRepository:   userLogRepository,
	}
	authHandler := &handler.AuthHandler{
		UnimplementedAuthServiceServer: unimplementedAuthServiceServer,
		TokenMiddleware:                tokenMiddleware,
		UserUseCase:                    userUseCase,
	}
	unimplementedUserServiceServer := pb.UnimplementedUserServiceServer{}
	userHandler := &handler.UserHandler{
		UnimplementedUserServiceServer: unimplementedUserServiceServer,
		UserUseCase:                    userUseCase,
		TokenMiddleware:                tokenMiddleware,
	}
	grpcProvider := &GRPCProvider{
		Conf:            conf,
		TokenMiddleware: tokenMiddleware,
		RoutesServices:  grpcMethodService,
		AuthHandler:     authHandler,
		UserHandler:     userHandler,
	}

	return grpcProvider
}
