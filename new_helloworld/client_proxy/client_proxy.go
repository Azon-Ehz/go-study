package client_proxy

import (
	"awesomeProject/new_helloworld/hanlder"
	"net/rpc"
)

type HelloServiceStub struct {

	//conn 的类型是 *rpc.Client
	//HelloServiceStub 有一个匿名字段 *rpc.Client
	*rpc.Client
}

// protocol 请求协议  address 地址
func NewHelloServiceClient(protocol string, address string) HelloServiceStub {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		panic("connect fail")
	}
	return HelloServiceStub{conn} //语法糖 因为类型和匿名字段类型一直 所有直接初始化匿名字段

}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(hanlder.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
