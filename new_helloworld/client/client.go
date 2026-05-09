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
}
