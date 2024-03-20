package main

import (
	"fmt"
	"log"
	"net/rpc"

	rpcserver "GolangLocalProject/go_advance_program/rpc_server"
)

type HelloServiceClient struct {
	*rpc.Client
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(rpcserver.HelloServiceName+".Hello", request, reply)
}

var _ rpcserver.HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func main() {
	client, err := DialHelloService("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	var reply string
	err = client.Hello("dfzj", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reply: %v\n", reply)
}
