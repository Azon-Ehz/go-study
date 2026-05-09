package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan int        //环形数组
	msg = make(chan int, 3) //无缓冲
	go func(msg chan int) { //go有一种happen-before的机制，可以保障
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("ALL DONE")
	}(msg)

	msg <- 1
	msg <- 2
	msg <- 3
	close(msg)
	msg <- 4 //panic 已经关闭的channel不能取值 但可以取值
	time.Sleep(3 * time.Second)

}
