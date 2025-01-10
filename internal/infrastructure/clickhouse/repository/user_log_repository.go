package repository

import (
	"context"
	"fmt"
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/magomedcoder/gskeleton/internal/infrastructure/clickhouse/model"
	"github.com/magomedcoder/gskeleton/pkg/clickhouseutil"
)

type IUserLogRepository interface {
	Create(ctx context.Context, userLog *model.UserLog) error
}

var _ IUserLogRepository = (*UserLogRepository)(nil)

type UserLogRepository struct {
	clickhouseutil.Repo[model.UserLog]
}

func NewUserLogRepository(clickHouse *clickHouseDriver.Conn) *UserLogRepository {
	return &UserLogRepository{Repo: clickhouseutil.NewRepo[model.UserLog](clickHouse)}
}

func (u *UserLogRepository) Create(ctx context.Context, userLog *model.UserLog) error {
	if err := u.Repo.Create(ctx, userLog); err != nil {
		return fmt.Errorf("не удалось записать лог: %s", err)
	}

	return nil
}
