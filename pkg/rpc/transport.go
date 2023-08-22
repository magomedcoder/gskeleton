package rpc

import (
	"context"
)

type Transport interface {
	Run(ctx context.Context) error
}
