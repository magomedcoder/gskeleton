package di

import (
	"github.com/magomedcoder/gskeleton/internal/cli/handler"
	"github.com/magomedcoder/gskeleton/internal/config"
	clickhouseRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/clickhouse/repository"
	postgresRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/repository"
	redisRepo "github.com/magomedcoder/gskeleton/internal/infrastructure/redis/repository"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/usecase"
)

type CLIProvider struct {
	Conf    *config.Config
	Migrate *handler.Migrate
}

func NewCliInjector(conf *config.Config) *CLIProvider {
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
	migrate := &handler.Migrate{
		Conf:        conf,
		Db:          db,
		UserUseCase: userUseCase,
	}
	cliProvider := &CLIProvider{
		Conf:    conf,
		Migrate: migrate,
	}

	return cliProvider
}
