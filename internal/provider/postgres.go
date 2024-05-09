package provider

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.Postgresql.User,
		config.Postgresql.Password,
		config.Postgresql.Host,
		config.Postgresql.Port,
		config.Postgresql.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("ошибка базы данных: %v", err))
	}

	return db
}
