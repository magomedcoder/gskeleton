package repository

import (
	"context"
	"encoding/json"
	"github.com/magomedcoder/gskeleton/internal/infrastructure/redis/model"
	"github.com/redis/go-redis/v9"
	"time"
)

type IUserCacheRepository interface {
	Set(ctx context.Context, key string, value model.UserCache, ttl int64) error

	Get(ctx context.Context, key string) (*model.UserCache, error)
}

var _ IUserCacheRepository = (*UserCacheRepository)(nil)

type UserCacheRepository struct {
	client *redis.Client
}

func NewUserCacheRepository(client *redis.Client) *UserCacheRepository {
	return &UserCacheRepository{client: client}
}

func (u *UserCacheRepository) Set(ctx context.Context, key string, value model.UserCache, ttl int64) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return u.client.Set(ctx, key, data, time.Duration(ttl)*time.Second).Err()
}

func (u *UserCacheRepository) Get(ctx context.Context, key string) (*model.UserCache, error) {
	data, err := u.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var value *model.UserCache
	if err := json.Unmarshal([]byte(data), &value); err != nil {
		return nil, err
	}

	return value, nil
}
