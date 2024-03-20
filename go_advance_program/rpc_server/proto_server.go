package rpcserver

import "GolangLocalProject/go_advance_program/hello"

type HelloProtoService struct{}

func (p *HelloProtoService) Hello(request *hello.String, reply *hello.String) error {
	reply.Value = "hello : " + request.GetValue()
	return nil
}
