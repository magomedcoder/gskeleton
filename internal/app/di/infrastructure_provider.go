package di

import (
	clickhouseRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/clickhouse/repository"
	postgresRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/repository"
	redisRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/redis/repository"
)

type InfrastructureProvider struct {
	UserCacheRepository *redisRepo.UserCacheRepository
	UserRepository      *postgresRepo.UserRepository
	UserLogRepository   *clickhouseRepo.UserLogRepository
}

func NewInfrastructureProvider(provider *Provider) *InfrastructureProvider {
	userCacheRepository := redisRepo.NewUserCacheRepository(provider.Redis)
	userRepository := postgresRepo.NewUserRepository(provider.Postgres)
	userLogRepository := clickhouseRepo.NewUserLogRepository(provider.ClickHouse)

	return &InfrastructureProvider{
		UserCacheRepository: userCacheRepository,
		UserRepository:      userRepository,
		UserLogRepository:   userLogRepository,
	}
}
