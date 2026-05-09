package server_proxy

import (
	"awesomeProject/new_helloworld/hanlder"
	"net/rpc"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv hanlder.HelloService) error {
	return rpc.RegisterName(hanlder.HelloServiceName, srv)
}
