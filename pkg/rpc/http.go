package rpc

import (
	"context"
	"log"
	"net"
	"net/http"
)

type HTTP struct{}

func (h *HTTP) Run(ctx context.Context, resolver Resolver) error {
	srv := http.Server{
		Addr: ":8000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
