package di

import (
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Provider struct {
	Redis      *redis.Client
	Postgres   *gorm.DB
	ClickHouse *clickHouseDriver.Conn
}

func NewProvider(conf *config.Config) *Provider {
	redisProvider := provider.NewRedisClient(conf)
	postgresProvider := provider.NewPostgresClient(conf)
	clickHouseProvider := provider.NewClickHouseClient(conf)

	return &Provider{
		Redis:      redisProvider,
		Postgres:   postgresProvider,
		ClickHouse: clickHouseProvider,
	}
}
