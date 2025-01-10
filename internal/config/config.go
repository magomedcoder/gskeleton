package config

import (
	"fmt"
	"github.com/magomedcoder/gskeleton/pkg/encrypt"
	"github.com/magomedcoder/gskeleton/pkg/strutil"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Server     *Server     `yaml:"server"`
	Postgres   *Postgres   `yaml:"postgres"`
	Redis      *Redis      `yaml:"redis"`
	ClickHouse *ClickHouse `yaml:"clickhouse"`
	Jwt        *Jwt        `yaml:"jwt"`
	sid        string
}

func New(filename string) *Config {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var conf Config
	if yaml.Unmarshal(content, &conf) != nil {
		panic(fmt.Sprintf("%s: %v", filename, err))
	}

	conf.sid = encrypt.Md5(fmt.Sprintf("%d%s", time.Now().UnixNano(), strutil.Random(6)))

	return &conf
}
