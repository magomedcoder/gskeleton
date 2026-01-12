package middleware

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

type LoggingMiddleware struct {
}

func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (l *LoggingMiddleware) UnaryLoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)

	log.Printf("Request - Method: %s \t Duration: %s \t Error: %v \n", info.FullMethod, time.Since(start), err)

	return h, err
}

func (l *LoggingMiddleware) StreamLoggingInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	err := handler(srv, ss)

	log.Printf("Stream - Method: %s \t Duration: %s \t Error: %v \n", info.FullMethod, time.Since(start), err)

	return err
}
