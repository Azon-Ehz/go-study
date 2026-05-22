package main

import (
	"awesomeProject/grpc_validate_test/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

type Validate interface {
	Validate() error
}

func (s *Server) SayHello(ctx context.Context, request *proto.Person) (*proto.Person, error) {
	return &proto.Person{
		Id: request.Id,
	}, nil
}

func main() {

	//拦截器
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		fmt.Println("接收到了一个新的请求")
		//grpc metaData验证器
		//类型转换成Validate
		if r, ok := req.(Validate); ok {
			if err = r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
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
