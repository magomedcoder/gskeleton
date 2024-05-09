package provider

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewPostgresDB(conf *config.Config) *gorm.DB {
	gormConfig := &gorm.Config{}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: conf.Postgres.GetDsn(),
	}), gormConfig)
	if err != nil {
		panic(fmt.Errorf("ошибка подключения к postgres: %v", err))
	}

	if db.Error != nil {
		panic(fmt.Errorf("ошибка базы данных: %v", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
