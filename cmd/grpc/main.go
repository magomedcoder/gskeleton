package main

import (
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/provider"
	_userRepo "github.com/magomedcoder/gskeleton/internal/repository/user/repo"
	_userService "github.com/magomedcoder/gskeleton/internal/service/user"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/handler"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/router"
	"log"
)

func main() {
	conf, err := config.ReadConfig("./configs/main.yaml")
	if err != nil {
		log.Fatalf("%v", err)
	}

	//srv, _ := Initialize(conf)
	//if err != nil {
	//	os.Exit(2)
	//}
	//
	//srv.Server.Start()
	db := provider.NewPostgresDB(conf)

	userRepository := _userRepo.NewUserRepository(db)
	userService := _userService.NewUserService(userRepository)

	tokenMiddleware := middleware.NewTokenMiddleware(conf)
	authMiddleware := middleware.NewAuthMiddleware(tokenMiddleware)

	authHandler := handler.NewAuthHandler(authMiddleware, tokenMiddleware, userService)
	userHandler := handler.NewUserHandler(authMiddleware, userService)

	routesServices := router.NewGrpMethodsService()

	srv := provider.NewGrpcServer(conf, authHandler, userHandler, authMiddleware, routesServices)

	srv.Start()
}
