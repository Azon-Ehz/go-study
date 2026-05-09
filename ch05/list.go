package main

import (
	"container/list"
	"fmt"
)

func main() {
	//List集合 两种定义方式
	//var myList list.List
	myList := list.New()
	myList.PushBack("go")
	myList.PushBack("grpc")
	myList.PushBack("gin")
	fmt.Println(myList) //0x00 地址
	myList.PushFront("PHP")
	myList.PushFront("Java")
	//循环遍历数据
	i := myList.Front()
	for ; i != nil; i = i.Next() {
		if i.Next().Value.(string) == "PHP" {
			break
		}
		fmt.Println(i.Value) //0x00 地址
	}
	myList.InsertBefore("Python", i)
	myList.Remove(i)
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value) //0x00 地址
	}

	//集合类型(复杂数据类型)一共四种
	//1.数组 - 不同长度的数组类型不一样 2.切片 - 动态数组 用起来方便 性能高 尽量使用 3.map - key-value键值对 4.list - 用的少
}
