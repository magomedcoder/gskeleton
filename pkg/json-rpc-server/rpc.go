package json_rpc_server

import (
	"context"
	"encoding/json"
)

type Request struct {
	JsonRpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	Id      any             `json:"id"`
}

type Response struct {
	JsonRpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   error           `json:"error,omitempty"`
	Id      any             `json:"id,omitempty"`
}

type Handler func(ctx context.Context, req *Request) *Response

type HandlerFunc func(context.Context, json.RawMessage) (json.RawMessage, error)

type Flusher interface {
	Flush()
}

func Param[Request any, Response any](handler func(context.Context, *Request) (Response, error)) HandlerFunc {
	return func(ctx context.Context, in json.RawMessage) (json.RawMessage, error) {
		req := new(Request)
		if err := json.Unmarshal(in, req); err != nil {
			return nil, ErrorFromCode(ErrCodeParseError)
		}
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		return json.Marshal(resp)
	}
}

func EmptyParam[Response any](handler func(context.Context) (Response, error)) HandlerFunc {
	return func(ctx context.Context, in json.RawMessage) (json.RawMessage, error) {
		resp, err := handler(ctx)
		if err != nil {
			return nil, err
		}

		return json.Marshal(resp)
	}
}
