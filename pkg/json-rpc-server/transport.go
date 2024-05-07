package json_rpc_server

import (
	"context"
	"io"
)

type Transport interface {
	Run(ctx context.Context, resolver Resolver) error
}

type Resolver interface {
	Resolve(ctx context.Context, reader io.Reader, writer io.Writer)
}
