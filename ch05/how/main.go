package main

import (
	"fmt"
	"unsafe"
)

func printSlice(data []string) {
}

type slice struct { //切片的底层体现
	array unsafe.Pointer //用来存储实际数据的指针，指向一块连续的内存
	len   int            //切片中元素的长度
	cap   int            //array 数组的长度
}

func main() {
	//go的slice 在函数参数传递的时候是值传递还是引用传递 ：：是值传递 效果又呈现出引用传递(不完全是)
	// 原始切片
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1 // 复制结构体，但指针指向同一数组

	// 情况1：修改元素 - 相互影响
	s2[0] = 100
	fmt.Println(s1[0]) // 输出 100，s1也被修改

	// 情况2：append导致扩容 - 不再共享
	s2 = append(s2, 7, 8, 9, 10) // 超出容量，创建新数组
	s2[2] = 300
	fmt.Println(s1[2]) // 输出 3，s1未受影响
	/*
			简单总结一下 slice 传递值还是引用的的底层问题
			1.每个slice实际上都是一个类似如下的函数
						type slice struct {
							array unsafe.Pointer //用来存储实际数据的指针，指向一块连续的内存
							len   int            //长度
							cap   int            //容量
						}
			2.当你在传递一个切片时，实际上是传递的这个切片的副本，由于他们现在指针仍然是同一个，所以是共享底层数据的
						所以当你在修改已有元素时会相互影响

		专业技术描述是
		1，没有真正的引用传递 只有值传递
		2.扩容是分界线 扩容之前共享数组 扩容后彻底分离
		3.copy make等显式操作会创建独立的数组空间
	*/
}
