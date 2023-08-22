package rpc

import (
	"context"
	"encoding/json"
)

type RpcRequest struct {
	Jsonrpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	Id      any             `json:"id"`
}

type RpcHandler func(ctx context.Context, req *RpcRequest)

type HandlerFunc func()
