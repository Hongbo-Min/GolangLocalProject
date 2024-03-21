package main

import (
	"context"
	"io"
	"log"
	"net"

	hello "GolangLocalProject/go_advance_program/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *hello.String) (*hello.String, error) {
	reply := &hello.String{
		Value: "hello: " + args.GetValue(),
	}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream hello.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		reply := &hello.String{Value: "hello: " + args.GetValue()}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {
	creds, err := credentials.NewServerTLSFromFile("./openssl_server/server.crt", "./openssl_server/server.key")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	hello.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(listener)
}
