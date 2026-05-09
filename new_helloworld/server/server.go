package main

import (
	"awesomeProject/new_helloworld/hanlder"
	"awesomeProject/new_helloworld/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	//编写服务三部曲
	//1. 实例化server
	listener, _ := net.Listen("tcp", ":1234")
	//2. 注册函数方法
	_ = server_proxy.RegisterHelloService(hanlder.HelloService{})
	//3. 启动服务
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
