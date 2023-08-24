package rpc

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

func Param[Response any](handler func(context.Context) (Response, error)) HandlerFunc {
	return func(ctx context.Context, in json.RawMessage) (json.RawMessage, error) {
		resp, err := handler(ctx)

		if err != nil {
			return nil, Error{
				Code:    ErrUser,
				Message: err.Error(),
			}
		}

		return json.Marshal(resp)
	}
}
