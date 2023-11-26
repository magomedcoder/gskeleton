package server

import (
	"context"
	"golang-app-skeleton/internal/config"
	"log"
	"net"
	"net/http"
)

type HTTP struct {
	App      config.App
	Uploader http.Handler
}

func (h *HTTP) Run(ctx context.Context, resolver Resolver) error {
	srv := http.Server{
		Addr: ":" + h.App.Port,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	http.HandleFunc("/json-rpc", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			resolver.Resolve(ctx, r.Body, w)
		} else {
			http.NotFound(w, r)
		}
	})

	if h.Uploader != nil {
		http.Handle("/upload", h.Uploader)
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
