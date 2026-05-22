package main

import (
	"awesomeProject/grpc_interpretor/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type customCredentials struct {
}

func (c customCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}
func (c customCredentials) RequireTransportSecurity() bool {
	return false
}

func main() {
	//interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	md := metadata.New(map[string]string{
	//		"appid":  "101010",
	//		"appkey": "i am key",
	//	})
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//
	//	start := time.Now()
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	fmt.Printf("耗时%s", time.Since(start))
	//	return err
	//}
	grpc.WithPerRPCCredentials(customCredentials{})
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(customCredentials{}))
	//建立拨号 发起连接
	conn, err := grpc.Dial("127.0.0.1:50051", opts...)
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
