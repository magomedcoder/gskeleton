package handler

import (
	"context"
	"github.com/magomedcoder/gskeleton"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/pkg/migrateutil"
	"gorm.io/gorm"
	"log"
)

type Migrate struct {
	Conf *config.Config
	Db   *gorm.DB
}

func (m *Migrate) Migrate(ctx context.Context) error {
	conn, err := m.Db.DB()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer conn.Close()

	migrator := migrateutil.MustGetNewMigrator(gskeleton.Migration(), "migrations/postgres")
	if err = migrator.ApplyMigrations(conn); err != nil {
		log.Fatalf("Ошибка при применении миграций: %v", err)
	}

	return nil
}
