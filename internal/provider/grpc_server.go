package provider

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"github.com/magomedcoder/gskeleton/internal/config"
	"github.com/magomedcoder/gskeleton/internal/transport/grpc/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

type grpcServer struct {
	config      *config.Config
	userHandler pb.UserServiceServer
	authHandler pb.AuthServiceServer
}

type Server interface {
	Serve() error
}

func NewGrpcServer(c *config.Config, us pb.UserServiceServer, ls pb.AuthServiceServer) Server {
	return &grpcServer{config: c, userHandler: us, authHandler: ls}
}

func (s *grpcServer) Serve() error {
	listener, err := net.Listen(
		s.config.Server.Grpc.GrpcProtocol,
		fmt.Sprintf("%s:%s", s.config.Server.Grpc.Host, s.config.Server.Grpc.GrpcPort),
	)
	if err != nil {
		return err
	}

	grpcLog := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)

	srv := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		middleware.LoggingServerInterceptor,
		middleware.AuthorizationServerInterceptor,
	)))

	pb.RegisterUserServiceServer(srv, s.userHandler)
	pb.RegisterAuthServiceServer(srv, s.authHandler)

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		return err
	}
	return nil
}
