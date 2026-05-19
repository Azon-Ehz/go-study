package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/rpc"

	"google.golang.org/grpc"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloRequest, error) {
	//err类型的返回值方法 Http的返回可以通过修改指针来进行操作

	return &proto.HelloRequest{
		Message: "Hello " + request.Name,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("接收到了一个新的请求")
		return handler(ctx, req)
	}

	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	proto.registerGreeterServer(g, new(Server))
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	err := g.Serve(lis)
	if err != nil {
		panic("failed to serve: " + err.Error())
	}
}
