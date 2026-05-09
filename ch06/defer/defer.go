package main

import "fmt"

func testDefer() (ret int) {
	defer func() {
		ret++ //defer 是有能力去修改你的返回值的
	}()
	return 10
}

func main() {
	//defer 语句的核心是在函数return之前执行 多个defer语句后进先出 也就是栈结构 压入与弹出数据
	ret := testDefer()
	fmt.Printf("ret is %d \r\n", ret)
}
