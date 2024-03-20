package main

import (
	"log"
	"net"
	"net/rpc"

	rpcserver "GolangLocalProject/go_advance_program/rpc_server"
)

func main() {
	err := rpcserver.RegisterHelloService(new(rpcserver.HelloService))
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}
}
