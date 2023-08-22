package rpc

import "net/http"

type HTTP struct{}

func (h *HTTP) Run() error {
	srv := http.Server{Addr: ":8000"}
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}
