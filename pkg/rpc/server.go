package rpc

import (
	"context"
	"encoding/json"
	"golang.org/x/sync/errgroup"
	"io"
	"strings"
)

type Server struct {
	transports []Transport
	handlers   map[string]HandlerFunc
}

func New(opts ...Option) *Server {
	server := &Server{
		transports: []Transport{},
		handlers:   map[string]HandlerFunc{},
	}
	for _, opt := range opts {
		opt(server)
	}

	return server
}

func (r *Server) Register(method string, handler HandlerFunc) {
	r.handlers[strings.ToLower(method)] = handler
}

func (r *Server) callMethod(ctx context.Context, req *RpcRequest) {
	h, ok := r.handlers[strings.ToLower(req.Method)]
	if !ok {

	}
	err := h(ctx, req.Params)
	if err != nil {

	}
}

func (r *Server) Resolve(ctx context.Context, rd io.Reader) {
	dec := json.NewDecoder(rd)
	for {
		req := new(RpcRequest)
		if err := dec.Decode(req); err != nil {
			break
		}
		exec := func() {
			h := r.callMethod
			h(ctx, req)
		}
		exec()
	}
}

func (r *Server) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, t := range r.transports {
		eg.Go(func(t Transport) func() error {
			return func() error {
				return t.Run(ctx, r)
			}
		}(t))
	}

	return eg.Wait()
}
