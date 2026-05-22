package main

import (
	"awesomeProject/grpc_token_auth_test/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}

func main() {
	//拦截器
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("接收到了一个新的请求")

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			fmt.Println("get metadata failed")
			return resp, status.Error(codes.Unauthenticated, "no token")
		}
		var (
			appid  string
			appkey string
		)
		if va1, ok := md["appid"]; ok {
			appid = va1[0]
		}
		if va1, ok := md["appkey"]; ok {
			appkey = va1[0]
		}
		if appid != "1010102" || appkey != "i am key" {
			return resp, status.Error(codes.Unauthenticated, "no token")
		}

		res, err := handler(ctx, req)
		fmt.Println("请求已完成")
		return res, err
	}

	opt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	err2 := g.Serve(lis)
	if err2 != nil {
		panic("failed to serve: " + err2.Error())
	}
}
