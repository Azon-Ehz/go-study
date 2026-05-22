package main

import (
	"awesomeProject/grpc_error_test/proto"
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	time.Sleep(time.Second * 5)
	return nil, status.Errorf(codes.NotFound, "记录不存在%s", request.Name)

	//err类型的返回值方法 Http的返回可以通过修改指针来进行操作
	//return &proto.HelloReply{
	//	Message: "Hello " + request.Name,
	//}, nil
}

func main() {

	/**
	server端 最好是帮我生成接口 我们只需要去每个接口中实现对应的业务逻辑
	客户端需要帮我们生成对应的方法 同时将这个方法绑定到一个结构体上，生成的时候 我们可能需要传参数
	比如 ip:port端口
	*/
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to serve: " + err.Error())
	}
}
