package main

import (
	"context"
	"fmt"
	"io"
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
	stream, err := client.Subscribe(context.Background(), &grpcpubsub.String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("reply.GetValue(): %v\n", reply.GetValue())
	}
}
