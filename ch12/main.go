package main

import (
	"fmt"
	"time"
)

//并发编程 go语言并发的特点的 简单 易写出高并发
//其他语言的并发编程 大部分都是通过多线程实现 特殊情况还会有多进程

//缺点
//多线程和多线程主要的问题是因为耗费内存 还会被操作系统调度
//内存占用很大 线程切换 比如php这些语言 启动线程 实际上是交给操作系统去做
//在java中 则是交予jvm jvm本质上也是交给操作系统
//web2.0后 对并发要求越来越高 多线程和多进程较为吃力 后面又出现了用户级线程 也叫绿程 有些地方叫轻量级线程 叫协程asyncio-python swoole-php netty-java
//协程最大的特点是内存占用小 切换快 go语言内存占比(2k)其他语言可能是(2mb)
//go语言没有线程 只有协程可用 所有第三方库全部支持协程-goroutine

func asyncPrint() {

}

// 主死随从
// 主协程 主协程如果挂了 子协程也挂
func main() {
	//匿名函数启动 goroutine
	//1. 闭包 2.for循环的坑
	for i := 0; i < 100; i++ {
		//goroutine 的生成与执行是没有顺序的
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main - goroutine")
	time.Sleep(5 * time.Second)

	//软件开发中 当有一个问题很难解决时 加一层 都可以解决

	//不同语言的线程调度模型不同
	//go GMP
	//在操作系统中申请线程 比如和cpu数量一样
	//假如你有10000个goroutine go把这些goroutine交给调度器(类似队列) 多对多 多个goroutine由调度器交给多个线程 M:N 线程可用解决多个goroutine的执行
	// 不存在把申请的线程撑爆 因为线程数是固定的 且调度器在调度时会等待
	//线程 M - thread线程 N- G goroutine

	//多个调度器 多个goroutine交给调度器 调度器找寻空闲thead线程 绑在一起执行 每个处理器之间有锁？
	//调度器中的goroutine会有序执行 如果其中某个goroutine中有sleep或者网络请求 会把当前g先挂起 执行后一个goroutine
}
