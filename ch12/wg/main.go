package main

import (
	"fmt"
	"sync"
)

// 子goroutine 如何通知 主goroutine自己结束了 主的goroutine如何知道子的goroutine 已经结束了

func main() {
	var wg sync.WaitGroup

	wg.Add(100)
	//我要监控多少个 goroutine 执行结束
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done() //依次-1注册的监控 - 100-1
			fmt.Println(i)
		}(i)
	}
	//等到所有子goroutine执行结束
	wg.Wait()
	fmt.Println("goroutine all down")
	//waitGroup主要用于等待goroutine执行结束
	//wg.add要与wg.down 成对出现 可以在goroutine中先声明 defer wg.down
	//wg.wait 等待所有携程全部结束
}
