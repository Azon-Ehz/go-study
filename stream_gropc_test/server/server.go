package main

import (
	"awesomeProject/stream_gropc_test/proto"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const PORT = ":50052"

type server struct {
	proto.UnimplementedGreeterServer
}

// 一元调用
//
//	func (s *server) GetStream(req *proto.StreamReqData) (*proto.StreamRspData, error) {
//		return nil, nil
//	}
//

// 服务端流模式
func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v %d", time.Now().Unix(), i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
		i++
	}
	return nil
}

// 客户端流模式
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		if r, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(r)
		}

	}
	return nil
}

// 双向流模式
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		allCount := 0
		for {
			allCount++
			fmt.Println("向客户端发送")
			allStr.Send(&proto.StreamResData{
				Data: "我是服务端 推流" + strconv.Itoa(allCount),
			})
		}
	}()

	go func() {
		defer wg.Done()
		for {
			allRep, err := allStr.Recv()
			if err != nil {
				fmt.Println(err)
				break
			} else {
				fmt.Println(allRep.Data)
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
