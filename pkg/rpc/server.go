package rpc

import (
	"context"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	transports []Transport
}

func New(opts ...Option) *Server {
	server := &Server{
		transports: []Transport{},
	}
	for _, opt := range opts {
		opt(server)
	}

	return server
}

func (r *Server) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, t := range r.transports {
		eg.Go(func(t Transport) func() error {
			return func() error {
				return t.Run(ctx)
			}
		}(t))
	}

	return eg.Wait()
}
