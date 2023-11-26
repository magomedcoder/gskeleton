package provider

import (
	"golang-app-skeleton/internal/config"
	"golang-app-skeleton/internal/transport/app/handler"
	"golang-app-skeleton/internal/transport/app/router"
	"golang-app-skeleton/pkg/server"
	"net/http"
)

func NewRpcServer(conf *config.Config, handler *handler.Handler) *server.Server {
	_http := &server.HTTP{
		App: conf.App,
	}

	_http.Uploader = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO Обработайте логику загрузки файла здесь
	})

	s := server.New(
		server.WithTransport(_http),
	)
	s = router.Methods(s, handler)

	return s
}
