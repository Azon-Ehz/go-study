package main

import (
	"fmt"
	"net/rpc"
	"time"

	"google.golang.org/grpc/metadata"
)

func main() {
	//建立拨号 发起连接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败")
	}

	md := metadata.Pairs("timestamp", time.Now().Format(timestamp))
	md := metadata.New(map[string]string{
		"timestamp": time.Now().Format(timestamp),
	})
	ctx := metadata.NewOutgoingContext(context.Fr)
	//var reply *string = new(string) //指针变量默认nil  new初始化创建地址
	var reply string //初始化默认值 有地址
	err = client.Call("HelloService.Hello", "Zinon", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
