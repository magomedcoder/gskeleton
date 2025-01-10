package config

import (
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
)

type ClickHouse struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (c *ClickHouse) Options() *clickhouse.Options {
	return &clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", c.Host, c.Port)},
		Auth: clickhouse.Auth{
			Database: c.Database,
			Username: c.Username,
			Password: c.Password,
		},
	}
}
