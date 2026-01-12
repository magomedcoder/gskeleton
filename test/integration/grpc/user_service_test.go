package grpc

import (
	"context"
	"fmt"
	"github.com/magomedcoder/gskeleton/api/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"
)

func TestAuthLoginService(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC серверу: %v", err)
	}

	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)
	res, err := client.Login(ctx, &pb.LoginRequest{
		Username: "test",
		Password: "test",
	})
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса Login: %v", err)
		return
	}

	if res.AccessToken == "" {
		t.Error("Получен пустой токен, ожидался валидный токен")
	}

	fmt.Printf("Token: %s\n", res.AccessToken)
}
