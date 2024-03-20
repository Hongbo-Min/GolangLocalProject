package rpcserver

import "net/rpc"

const HelloServiceName = "HelloService"

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello : " + request
	return nil
}

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
