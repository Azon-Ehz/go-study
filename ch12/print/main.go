package main

import (
	"fmt"
	"time"
)

// 向一个无缓冲 channel 发送数据时，发送方会阻塞，直到另一个 goroutine 从这个 channel 接收数据；
// 反之，从一个无缓冲 channel 接收数据时，接收方也会阻塞
var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		<-number
		fmt.Printf("%d%d", i, i+1)
		letter <- true
		i += 2
	}
}

func printLetter() {

	str := "ABCDEFGHIJKLMNOPQRSTUVTXYZ"
	for i := 0; i < 24; i += 2 {
		<-letter
		if i >= len(str) {
			return
		}
		fmt.Printf("%s", str[i:i+2])
		number <- true
	}
}

func main() {

	go printNum()
	go printLetter()
	number <- true

	time.Sleep(time.Second * 2)
}
