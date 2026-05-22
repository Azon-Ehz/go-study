package main

import (
	"awesomeProject/helloworld/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	//建立拨号 发起连接
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "bobby"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
