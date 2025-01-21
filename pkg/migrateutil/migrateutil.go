package migrateutil

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"

	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/clickhouse"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type ConfPostgres struct {
	Username string
	Password string
	Host     string
	Port     int64
	Database string
}

type ConfClickHouse struct {
	Username string
	Password string
	Host     string
	Port     int64
	Database string
}

type Conf struct {
	Postgres   *ConfPostgres
	ClickHouse *ConfClickHouse
}

type Migrator struct {
	Migration embed.FS
}

func NewMigrator(
	migration embed.FS,
) *Migrator {
	return &Migrator{
		Migration: migration,
	}
}

func (m *Migrator) Postgres(db *sql.DB) error {

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("не удалось создать экземпляр базы данных: %v", err)
	}

	d, err := iofs.New(m.Migration, "migrations/postgresql")
	if err != nil {
		return fmt.Errorf("не удалось загрузить миграции: %v", err)
	}

	migrator, err := migrate.NewWithInstance("migration_embeded_sql_files", d, "psql_db", driver)
	if err != nil {
		return fmt.Errorf("не удалось создать миграцию: %v", err)
	}

	defer migrator.Close()

	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("не удалось применить миграции: %v", err)
	}

	return nil
}

func (m *Migrator) ClickHouse(db *sql.DB) error {
	driver, err := clickhouse.WithInstance(db, &clickhouse.Config{
		MultiStatementEnabled: true, // Включить выполнение нескольких запросов
	})
	if err != nil {
		log.Fatalf("Не удалось создать драйвер миграций: %v", err)
	}

	d, err := iofs.New(m.Migration, "migrations/clickhouse")
	if err != nil {
		return fmt.Errorf("не удалось загрузить миграции: %v", err)
	}

	migrator, err := migrate.NewWithInstance("migration_embeded_sql_files", d, "ch_db", driver)
	if err != nil {
		return fmt.Errorf("не удалось создать миграцию: %v", err)
	}

	defer migrator.Close()

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("не удалось применить миграции: %v", err)
	}

	return nil
}
