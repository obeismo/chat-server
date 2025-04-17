package main

import (
	"context"
	"log"
	"time"

	"github.com/fatih/color"
	desc "github.com/obeismo/chat_server/grpc/pkg/chat_server/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	usernames := []string{"lalala", "tralalelo", "tralala"}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %s", err)
	}
	defer conn.Close()

	c := desc.NewChatServerV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Create(ctx, &desc.CreateRequest{Usernames: usernames})
	if err != nil {
		log.Fatalf("failed to get user by id: %s", err)
	}

	log.Printf(color.RedString("User ID:\n"), color.GreenString("%+v", r.GetId()))
}
