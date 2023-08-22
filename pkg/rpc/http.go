package rpc

import (
	"context"
	"log"
	"net/http"
)

type HTTP struct{}

func (h *HTTP) Run(ctx context.Context) error {
	srv := http.Server{Addr: ":8000"}
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
