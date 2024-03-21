package main

import (
	"context"
	"log"

	grpcpubsub "GolangLocalProject/go_advance_program/rpc_pubsub"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := grpcpubsub.NewPubsubServiceClient(conn)
	_, err = client.Publish(context.Background(), &grpcpubsub.String{Value: "golang: hello go"})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(context.Background(), &grpcpubsub.String{Value: "docker: hello docker"})
	if err != nil {
		log.Fatal(err)
	}
}
