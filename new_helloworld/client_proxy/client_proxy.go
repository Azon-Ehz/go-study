package client_proxy

import (
	"awesomeProject/new_helloworld/hanlder"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol string, address string) HelloServiceStub {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		panic("connect fail")
	}
	return HelloServiceStub{conn}

}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(hanlder.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
