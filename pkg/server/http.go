package server

import (
	"context"
	"golang-app-skeleton/internal/config"
	"log"
	"net"
	"net/http"
)

type HTTP struct {
	App config.App
}

func (h *HTTP) Run(ctx context.Context, resolver Resolver) error {
	srv := http.Server{
		Addr: ":" + h.App.Port,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			resolver.Resolve(ctx, r.Body, w)
		}),
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		<-ctx.Done()
		srv.Close()
	}()

	log.Println("Запущен")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}
