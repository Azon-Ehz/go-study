package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var RWlock sync.RWMutex

func cpuInfoFunc(ctx context.Context) {
	//这里能拿到一个请求的id
	fmt.Printf("traceId:%s\r\n", ctx.Value("traceid")) //traceId:order_id
	//记录日志 这次日志是哪次请求打印的
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpuInfo")
		}
	}
}

func main() {
	wg.Add(1)
	//context提供了三个函数 WithCancel/WithTimeout/WithValue 本质上是三种不同类型的结构体
	//如果你的goroutine 函数中希望被控制，超时/传值，但我不希望影响原来的接口信息时，函数的第一个参数就要尽量加一个ctx
	//1.cancel
	//ctx1, cancel1 := context.WithCancel(context.Background())
	//ctx2, _ := context.WithCancel(ctx1)

	//2.timeout 主动超时
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	//3.withDeadline 在时间点cancel

	//4.withValue 主动传值
	valueCtx := context.WithValue(ctx, "traceid", "order_id")
	go cpuInfoFunc(valueCtx)
	wg.Wait()
	fmt.Println("监测完成")
}
