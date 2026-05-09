package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	//err类型的返回值方法 Http的返回可以通过修改指针来进行操作
	*reply = "Hello, " + request
	return nil
}

func main() {
	//编写服务三部曲
	_ = rpc.RegisterName("HelloService", &HelloService{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.ReadCloser
			io.Writer
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
}
