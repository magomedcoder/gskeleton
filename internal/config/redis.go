package config

import (
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Host string `yaml:"host"`
}

func (r *Redis) Options() *redis.Options {
	return &redis.Options{
		Addr:        r.Host,
		ReadTimeout: -1,
	}
}
