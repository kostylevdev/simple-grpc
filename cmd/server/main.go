package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/kostylevdev/simple-grpc/proto/notification"
	"google.golang.org/grpc"
)

type server struct {
	notification.UnimplementedNotificationServiceServer
}

func (s *server) Notify(ctx context.Context, req *notification.NotificationRequest) (*notification.NotificationResponse, error) {
	fmt.Printf("message: %s\n", req.Message)
	return &notification.NotificationResponse{
		Message: "from server to client",
		Status:  "ok",
	}, nil
}

func (s *server) mustEmbedUnimplementedNotificationServiceServer() {}
func main() {
	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	notification.RegisterNotificationServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
