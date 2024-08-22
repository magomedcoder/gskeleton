package provider

import (
	"context"
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/router"
	authPb "github.com/magomedcoder/gskeleton/pkg/pb_generated/auth"
	userPb "github.com/magomedcoder/gskeleton/pkg/pb_generated/user"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

type GrpcServer struct {
	Conf           *config.Config
	AuthHandler    authPb.AuthServiceServer
	UserHandler    userPb.UserServiceServer
	AuthMiddleware *middleware.AuthMiddleware
	RoutesServices *router.GrpcMethodService
}

type IServer interface {
	Serve() error
	Start()
}

func NewGrpcServer(
	conf *config.Config,
	authHandler authPb.AuthServiceServer,
	userHandler userPb.UserServiceServer,
	authMiddleware *middleware.AuthMiddleware,
	routesServices *router.GrpcMethodService,
) IServer {
	return &GrpcServer{
		Conf:           conf,
		AuthHandler:    authHandler,
		UserHandler:    userHandler,
		AuthMiddleware: authMiddleware,
		RoutesServices: routesServices,
	}
}

func (g *GrpcServer) Serve() error {
	listener, err := net.Listen(
		g.Conf.Server.Grpc.GrpcProtocol,
		fmt.Sprintf("%s:%s", g.Conf.Server.Grpc.Host, g.Conf.Server.Grpc.GrpcPort),
	)
	if err != nil {
		return err
	}

	grpcLog := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)

	srv := grpc.NewServer(grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		middleware.LoggingServerInterceptor,
		middleware.AuthorizationServerInterceptor,
	)))

	authPb.RegisterAuthServiceServer(srv, g.AuthHandler)
	userPb.RegisterUserServiceServer(srv, g.UserHandler)

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (g *GrpcServer) Start() {
	ctx := context.Background()

	ctx = middleware.RegisterGlobalService(ctx, g.AuthMiddleware)
	ctx = middleware.RegisterGlobalService(ctx, g.RoutesServices)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	group, _ := errgroup.WithContext(ctx)
	group.Go(func() error {
		log.Printf(
			"gRPC server running at %s://%s:%s \n",
			g.Conf.Server.Grpc.GrpcProtocol,
			g.Conf.Server.Grpc.Host,
			g.Conf.Server.Grpc.GrpcPort,
		)
		return g.Serve()
	})

	log.Fatal(group.Wait())
}
