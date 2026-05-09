package main

import (
	"fmt"
	"time"
)

func producr(out chan<- int) { // 写

	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)

}

func consumer(in <-chan int) { // 读
	for num := range in {
		fmt.Printf("num = %d\r\n", num)
	}
}

func main() {
	//默认情况下 channel 是双向的 既可以写也可以读
	//单向channel就是仅读或写
	//var ch1 chan int       // 双向
	//var ch2 chan<- float64 // 单向写
	//var ch3 <-chan float64 // 单向读
	//
	//c := make(chan int, 3)
	//var send chan<- int = c //sned-only
	//
	//var read <-chan int = c //recv-only
	////<-read
	////fmt.Println(c, send, read)
	//send <- 1
	//<-read
	c := make(chan int)
	go producr(c)
	go consumer(c)
	time.Sleep(10 * time.Second)
}
