package main

import (
	"context"
	"log"
	"net"

	"gRPCStudy/proto/greeter"

	"google.golang.org/grpc"
)

type server struct {
	greeter.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *greeter.HelloRequest) (rsp *greeter.HelloReply, err error) {
	rsp = &greeter.HelloReply{Message: "Hello " + req.Name}
	return rsp, nil
}

func main() {
	listener, err := net.Listen("tcp", ":52001")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	greeter.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
