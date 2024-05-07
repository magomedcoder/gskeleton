package provider

import (
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/transport/json-rpc/handler"
	"github.com/magomedcoder/gskeleton/internal/transport/json-rpc/router"
	"github.com/magomedcoder/gskeleton/pkg/json-rpc-server"
	"net/http"
)

func NewJsonRpcServer(conf *config.Config, handler *handler.Handler) *json_rpc_server.Server {
	_http := &json_rpc_server.HTTP{
		Addr: ":" + conf.Server.JsonRpc.Port,
	}

	_http.Uploader = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO Обработайте логику загрузки файла здесь
	})

	s := json_rpc_server.New(
		json_rpc_server.WithTransport(_http),
	)
	s = router.Methods(s, handler)

	return s
}
