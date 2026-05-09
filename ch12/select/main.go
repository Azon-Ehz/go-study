package main

import (
	"fmt"
	"sync"
	"time"
)

var done = make(chan struct{}) //channel 是多线程安全是 channel要初始化切记
var lock sync.Mutex

func g1(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	ch <- struct{}{}
}
func g2(ch chan struct{}) {
	time.Sleep(3 * time.Second)
	ch <- struct{}{}
}

func main() {
	//select 类似于 switch case 语句
	//select类似于 switch case语句，//select 主要作用于多个channel
	//但是select的功能和我们操作linux里面提供的io的select、poll、epoll
	//现在有个需求，
	//我们现在有两个goroutine都在执行，
	//但是我在主的goroutine中， 当某一个执行完成以后，这个时候我会立马知道
	g1Channel := make(chan struct{}, 1)
	g2Channel := make(chan struct{}, 2)
	//g1Channel <- struct{}{}
	//g2Channel <- struct{}{}
	go g2(g2Channel)
	go g1(g1Channel)

	//1.某一个分支就绪了就执行该分支 2.如果两个都就绪了先执行哪个？实际上都执行顺序是随机性的 目的是：防止饥饿
	//应用场景
	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-g1Channel:
			fmt.Println("g1 channel closed")
		case <-g2Channel:
			fmt.Println("g2 channel closed")
		case <-timer.C:
			fmt.Println("default")
			return
		}
	}
}
