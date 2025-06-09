package server

import (
	"context"
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/app/di"
	"github.com/magomedcoder/gskeleton/internal/delivery/grpc/middleware"
	cliV2 "github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func GRPC(ctx2 *cliV2.Context, app *di.GRPCProvider) error {
	ctx := context.Background()

	ctx = middleware.RegisterGlobalService(ctx, app.TokenMiddleware)
	ctx = middleware.RegisterGlobalService(ctx, app.RoutesServices)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	group, _ := errgroup.WithContext(ctx)
	group.Go(func() error {
		log.Printf(
			"gRPC server running at %s://%s:%d \n",
			app.Conf.Server.Grpc.GrpcProtocol,
			app.Conf.Server.Grpc.Host,
			app.Conf.Server.Grpc.GrpcPort,
		)
		listener, err := net.Listen(
			app.Conf.Server.Grpc.GrpcProtocol,
			fmt.Sprintf("%s:%d", app.Conf.Server.Grpc.Host, app.Conf.Server.Grpc.GrpcPort),
		)
		if err != nil {
			return err
		}

		grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))

		srv := grpc.NewServer(grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			middleware.LoggingServerInterceptor,
			middleware.AuthorizationServerInterceptor,
		)))

		pb.RegisterAuthServiceServer(srv, app.AuthHandler)
		pb.RegisterUserServiceServer(srv, app.UserHandler)

		reflection.Register(srv)

		if err := srv.Serve(listener); err != nil {
			return err
		}

		return nil
	})

	log.Fatal(group.Wait())

	return nil
}
