package main

import (
	"context"
	"log"

	"github.com/kostylevdev/simple-grpc/proto/notification"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect client: %s", err)
	}
	defer conn.Close()

	c := notification.NewNotificationServiceClient(conn)

	response, err := c.Notify(context.Background(), &notification.NotificationRequest{
		Message: "from client to server",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("message: %s, status: %s", response.Message, response.Status)
}
