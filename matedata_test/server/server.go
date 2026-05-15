package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	//err类型的返回值方法 Http的返回可以通过修改指针来进行操作
	*reply = "Hello, " + request
	return nil
}

func main() {
	//编写服务三部曲
	//1. 实例化server
	listener, _ := net.Listen("tcp", ":1234")
	//2. 注册函数方法
	_ = rpc.RegisterName("HelloService", &HelloService{})
	//3. 启动服务
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)
}
