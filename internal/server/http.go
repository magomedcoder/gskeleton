package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/magomedcoder/gskeleton/internal/app/di"
	cliV2 "github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HTTP(ctx *cliV2.Context, app *di.HTTPProvider) error {
	eg, groupCtx := errgroup.WithContext(context.Background())

	gin.SetMode(gin.ReleaseMode)

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Conf.Server.Http.Port),
		Handler: app.Engine,
	}

	eg.Go(func() error {
		log.Printf("Http server running at %d | PID %d", app.Conf.Server.Http.Port, os.Getpid())
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
