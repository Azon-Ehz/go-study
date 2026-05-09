package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	/*
			不要来通过共享内存来通信 而是要通过通信来共享内存
		 	其他编程语言在多线程编程时大部分都采用全局变量的方式
			也会提供消息队列的机制， 消费者和生产者之间的关系
			channel再加上语法糖 让channel的使用更加简单
	*/
	var msg chan string //环形数组

	/**
	有/无缓存channel的应用场景
	无缓存适用于 “通知” 比如事件 B需要第一时间知道A是否已完成
	有缓冲适用于 “消费/生产之间通信”
	1.消息传递、消息过滤
	2.信号广播
	3.事件订阅/广播
	4.任务分发
	5.结果汇总
	6.并发控制
	7.同步/异步
	上述都首选使用channel来完成

	*/

	//channel如果初始化值是0 那么放值进去会导致阻塞
	msg = make(chan string, 0) //无缓冲
	wg.Add(1)
	go func(msg chan string) { //go有一种happen-before的机制，可以保障
		data := <-msg
		fmt.Println(data)
		defer wg.Done()
	}(msg)
	msg <- "无缓存"
	wg.Wait()

	//有缓冲和无缓冲
	msg = make(chan string, 0)  //无缓冲队列
	msg = make(chan string, 10) //有缓冲队列
	msg <- "Zinon-"
	msg <- "Study-"
	msg <- "Golang"
	//data2 := <-msg //这样才是取值 并且是先进先出的去消费
	var build strings.Builder
	build.WriteString(<-msg) // deadlock 如果消费不存在的元素 那么也会造成deadlock
	build.WriteString(<-msg)
	build.WriteString(<-msg)
	data3 := build.String()
	//fmt.Println(data2)
	fmt.Println(data3)
}
