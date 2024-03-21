package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"GolangLocalProject/go_advance_program/pubsub"
	grpcpubsub "GolangLocalProject/go_advance_program/rpc_pubsub"

	"google.golang.org/grpc"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(ctx context.Context, arg *grpcpubsub.String) (*grpcpubsub.String, error) {
	p.pub.Publish(arg.GetValue())
	return &grpcpubsub.String{}, nil
}

func (p *PubsubService) Subscribe(arg *grpcpubsub.String, stream grpcpubsub.PubsubService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&grpcpubsub.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	grpcpubsub.RegisterPubsubServiceServer(grpcServer, new(PubsubService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(listener)
}
