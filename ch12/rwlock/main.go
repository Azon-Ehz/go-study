package main

import (
	"fmt"
	"sync"
	"time"
)

//锁本质上是将并行的代码 串行化了 非常影响性能
//即使是设计锁 那么也应该尽量的保证并行
// 我们有两组线程 两组携程 一组负责写 一组负责读
//绝大部份场景都是读多写少 读 > 写
//多个goroutine 读应该是并发 但读和写应该是串行 读和读之间也应该串行
//读写锁

func main() {
	var wg sync.WaitGroup
	var rwlock sync.RWMutex //读写锁
	wg.Add(6)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				rwlock.RLock()
				time.Sleep(500 * time.Millisecond)
				fmt.Println("get read lock") //读锁不会阻止其他读锁
				rwlock.RUnlock()
			}
		}()

		//写 write
		go func() {
			time.Sleep(1 * time.Second)
			defer wg.Done()
			rwlock.Lock() //加写锁 可以阻止别的写锁获取 和读锁的获取
			defer rwlock.Unlock()
			fmt.Println("get write lock")
			time.Sleep(5 * time.Second)
		}()
	}
	wg.Wait()
}
