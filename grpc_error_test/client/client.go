package main

import (
	"awesomeProject/grpc_error_test/proto"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	//建立拨号 发起连接
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	_, err = c.SayHello(ctx, &proto.HelloRequest{Name: "bobby"})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			fmt.Println("解析 error 失败")
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	//fmt.Println(r.Message)
}
