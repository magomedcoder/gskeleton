package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/magomedcoder/gskeleton"
	"github.com/magomedcoder/gskeleton/internal/config"
	postgresModel "github.com/magomedcoder/gskeleton/internal/infrastructure/postgres/model"
	"github.com/magomedcoder/gskeleton/internal/usecase"
	"github.com/magomedcoder/gskeleton/pkg/migrateutil"
	"gorm.io/gorm"
	"log"
	"time"
)

type Migrate struct {
	Conf        *config.Config
	Db          *gorm.DB
	UserUseCase usecase.IUserUseCase
}

func (m *Migrate) Migrate(ctx context.Context) error {
	postgres, err := sql.Open(
		"postgres", fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow",
			m.Conf.Postgres.Host,
			m.Conf.Postgres.Port,
			m.Conf.Postgres.Username,
			m.Conf.Postgres.Password,
			m.Conf.Postgres.Database,
		),
	)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	defer postgres.Close()

	clickHouse, err := sql.Open(
		"clickhouse",
		fmt.Sprintf(
			"clickhouse://%s:%s@%s:%d/%s",
			m.Conf.ClickHouse.Username,
			m.Conf.ClickHouse.Password,
			m.Conf.ClickHouse.Host,
			m.Conf.ClickHouse.Port,
			m.Conf.ClickHouse.Database,
		))
	if err != nil {
		log.Fatalf("не удалось подключиться к базе данных: %v", err)
	}

	defer clickHouse.Close()

	migrator := migrateutil.NewMigrator(gskeleton.Migration())
	if err := migrator.Postgres(postgres); err != nil {
		return errors.New(fmt.Sprintf("Ошибка при применении миграций Postgres: %v", err))
	}

	if err := migrator.ClickHouse(clickHouse); err != nil {
		return errors.New(fmt.Sprintf("Ошибка при применении миграций ClickHouse: %v", err))
	}

	return nil
}

func (m *Migrate) CreateUser(ctx context.Context) error {
	if _, err := m.UserUseCase.Create(ctx, &postgresModel.User{
		Username:  "admin",
		Password:  "admin123",
		CreatedAt: time.Now(),
	}); err != nil {
		return err
	}

	return nil
}
