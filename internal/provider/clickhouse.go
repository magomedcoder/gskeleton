package provider

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	clickHouseDriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/magomedcoder/gskeleton/internal/config"
)

func NewClickHouseClient(conf *config.Config) *clickHouseDriver.Conn {
	conn, err := clickhouse.Open(conf.ClickHouse.Options())
	if err != nil {
		panic(fmt.Errorf("ошибка подключения к ClickHouse: %v", err))
	}

	ctx := context.Background()
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		if err != nil {
			panic(fmt.Errorf("ошибка подключения к ClickHouse: %v", err))
		}
	}

	return &conn
}
