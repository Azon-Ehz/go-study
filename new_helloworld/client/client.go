package main

import (
	"awesomeProject/new_helloworld/client_proxy"
	"fmt"
)

func main() {
	//建立拨号 发起连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	var reply string //初始化默认值 有地址
	err := client.Hello("Zinon", &reply)
	if err != nil {
		panic("调用失败." + err.Error())
	}
	fmt.Println(reply)

	//1. 这个new_helloworld是grpc概念的印证
	//2. server_proxy和client_proxy都是可以自动生成的 而且是可以面向不同语言
	//3. 都可以满足 那就是 protobuf + grpc
}
