package main

import (
	"context"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/provider"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/router"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	ctx := context.Background()

	conf, err := config.ReadConfig("./configs/main.yaml")
	if err != nil {
		log.Fatalf("%v", err)
	}

	routesServices := router.NewGrpMethodsService()

	db, err := provider.NewPostgresDB(conf)
	if err != nil {

	}

	userRepository := InitializeUserRepository(db)

	tokenHandler := InitializeTokenHandler(conf.Jwt)

	userService := InitializeUserService(userRepository)
	authService := InitializeAuthService(tokenHandler)

	userHandler := InitializeUserHandler(userService, authService)
	authHandler := InitializeAuthHandler(authService, userService, tokenHandler)

	ctx = middleware.RegisterGlobalService(ctx, authService)
	ctx = middleware.RegisterGlobalService(ctx, routesServices)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		srv := provider.NewGrpcServer(conf, userHandler, authHandler)
		log.Printf(
			"gRPC json-rpc-server running at %s://%s:%s \n",
			conf.Server.Grpc.GrpcProtocol,
			conf.Server.Grpc.Host,
			conf.Server.Grpc.GrpcPort,
		)
		return srv.Serve()
	})
	log.Fatal(g.Wait())

}
