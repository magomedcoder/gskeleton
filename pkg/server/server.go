package server

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

func ResultResponse(id any, resp json.RawMessage) *Response {
	return &Response{
		JsonRpc: "2.0",
		Result:  resp,
		Id:      id,
	}
}

func ErrorResponse(id any, err error) *Response {
	return &Response{
		JsonRpc: "2.0",
		Error:   err,
		Id:      id,
	}
}

func (r *Server) callMethod(ctx context.Context, req *Request) *Response {
	h, ok := r.handlers[strings.ToLower(req.Method)]
	if !ok {
		return ErrorResponse(req.Id, ErrorFromCode(ErrCodeMethodNotFound))
	}
	resp, err := h(ctx, req.Params)
	if err != nil {
		return ErrorResponse(req.Id, err)
	}

	return ResultResponse(req.Id, resp)
}

func (r *Server) Resolve(ctx context.Context, rd io.Reader, w io.Writer) {
	dec := json.NewDecoder(rd)
	enc := json.NewEncoder(w)
	for {
		req := new(Request)
		if err := dec.Decode(req); err != nil {
			break
		}
		exec := func() {
			h := r.callMethod
			resp := h(ctx, req)
			if req.Id == nil {
				return
			}
			if err := enc.Encode(resp); err != nil {
				enc.Encode(ErrorResponse(req.Id, ErrorFromCode(ErrCodeInternalError)))
			}
			if w, canFlush := w.(Flusher); canFlush {
				w.Flush()
			}
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
