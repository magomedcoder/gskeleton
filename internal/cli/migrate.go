package cli

import (
	"github.com/magomedcoder/gskeleton"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/pkg/migrate"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"log"
)

type AppProvider struct {
	Conf *config.Config
	Db   *gorm.DB
}

func Migrate(ctx *cli.Context, app *AppProvider) error {
	conn, err := app.Db.DB()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer conn.Close()

	migrator := migrate.MustGetNewMigrator(gskeleton.Migration(), "migrations/postgres")
	if err = migrator.ApplyMigrations(conn); err != nil {
		log.Fatalf("Ошибка при применении миграций: %v", err)
	}

	return nil
}
