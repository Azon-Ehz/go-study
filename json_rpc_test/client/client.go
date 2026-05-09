package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//建立拨号 发起连接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败")
	}
	//var reply *string = new(string) //指针变量默认nil  new初始化创建地址
	var reply string //初始化默认值 有地址
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "Zinon2", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
