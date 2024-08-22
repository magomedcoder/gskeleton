package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	cliv2 "github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx *cliv2.Context, app *AppProvider) error {
	eg, groupCtx := errgroup.WithContext(context.Background())

	gin.SetMode(gin.ReleaseMode)

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8000),
		Handler: app.Engine,
	}

	eg.Go(func() error {
		log.Printf("Http server running at | PID %d", os.Getpid())
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		defer func() {
			log.Println("Выключение сервера...")
			timeCtx, timeCancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer timeCancel()
			if err := server.Shutdown(timeCtx); err != nil {
				log.Fatalf("Ошибка остановки сервера: %s", err)
			}
		}()
		select {
		case <-groupCtx.Done():
			return groupCtx.Err()
		case <-c:
			return nil
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("Принудительное завершение сервера: %s", err)
	}

	log.Println("Выход из сервера")

	return nil
}
