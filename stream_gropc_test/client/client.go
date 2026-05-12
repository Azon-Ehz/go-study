package main

import (
	"awesomeProject/stream_gropc_test/proto"
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	//服务端流模式
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "我要学习GO/GO/GO"})
	for {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}

	//服务端流模式
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("服务端流模式Go/GO/GO %d", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	//双向推流模式
	wg := sync.WaitGroup{}
	wg.Add(2)
	allS, _ := c.AllStream(context.Background())
	go func() {
		defer wg.Done()
		allCount := 0
		for {
			allCount++
			fmt.Println("向服务端发送")
			allS.Send(&proto.StreamReqData{
				Data: "我是客户端 推流" + strconv.Itoa(allCount),
			})
		}
	}()
	go func() {
		defer wg.Done()
		for {
			allRes, allErr := allS.Recv()
			if allErr != nil {
				fmt.Println(allErr)
				break
			} else {
				fmt.Println(allRes.Data)
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()

}
