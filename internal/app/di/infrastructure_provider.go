package di

import (
	clickhouseRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/clickhouse/repository"
	postgresRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/repository"
	redisRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/redis/repository"
)

type InfrastructureProvider struct {
	UserRepository          *postgresRepo.UserRepository
	UserCacheRepository     *redisRepo.UserCacheRepository
	JwtTokenCacheRepository *redisRepo.JwtTokenCacheRepository
	UserLogRepository       *clickhouseRepo.UserLogRepository
}

func NewInfrastructureProvider(provider *Provider) *InfrastructureProvider {
	userRepository := postgresRepo.NewUserRepository(provider.Postgres)
	userCacheRepository := redisRepo.NewUserCacheRepository(provider.Redis)
	jwtTokenCacheRepository := redisRepo.NewJwtTokenCacheRepository(provider.Redis)
	userLogRepository := clickhouseRepo.NewUserLogRepository(provider.ClickHouse)

	return &InfrastructureProvider{
		UserRepository:          userRepository,
		UserCacheRepository:     userCacheRepository,
		JwtTokenCacheRepository: jwtTokenCacheRepository,
		UserLogRepository:       userLogRepository,
	}
}
