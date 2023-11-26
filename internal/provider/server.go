package provider

import (
	"golang-app-skeleton/internal/config"
	"golang-app-skeleton/internal/transport/app/handler"
	"golang-app-skeleton/internal/transport/app/router"
	"golang-app-skeleton/pkg/server"
)

func NewRpcServer(conf *config.Config, handler *handler.Handler) *server.Server {
	http := &server.HTTP{
		App: conf.App,
	}
	server := server.New(
		server.WithTransport(http),
	)
	server = router.Methods(server, handler)

	return server
}
