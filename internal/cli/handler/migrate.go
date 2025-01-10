package handler

import (
	"context"
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
	db, err := m.Db.DB()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	defer db.Close()

	migrator := migrateutil.MustGetNewMigrator(gskeleton.Migration(), "migrations/postgres")
	if err = migrator.ApplyMigrations(db); err != nil {
		return errors.New(fmt.Sprintf("Ошибка при применении миграций: %v", err))
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
