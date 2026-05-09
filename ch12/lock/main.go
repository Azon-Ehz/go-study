package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
锁的目的是对共享资源竞争
*/
var total int32

var wg sync.WaitGroup

var lock sync.Mutex

//锁不可以复制 否则会导致失去锁的效果

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&total, 1)
		//lock.Lock()
		//total++
		//lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&total, -1)
		//lock.Lock()
		//total--
		//lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println("all goroutines down")
	fmt.Println(total)
}
