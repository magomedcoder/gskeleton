package server

import (
	"context"
	"fmt"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/app/di"
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

		srv := grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				app.Middleware.Logging.UnaryLoggingInterceptor,
				app.Middleware.Auth.UnaryAuthInterceptor,
			),
			grpc.ChainStreamInterceptor(
				app.Middleware.Logging.StreamLoggingInterceptor,
				app.Middleware.Auth.StreamAuthInterceptor,
			),
		)

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
