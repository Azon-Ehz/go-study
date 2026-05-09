package main

import (
	"fmt"
	"sync"
)

func main() {
	//map是一个key-val的无序集合 时间复杂度o(1)
	var courseMap = map[string]string{
		"go":   "golang",
		"grpc": "grpc入门",
		"gin":  "gin深入理解",
	}
	courseMap["name"] = "Zinon"
	//map是无序的 每次打印不保证都是完全一致的顺序
	for key, value := range courseMap {
		fmt.Println(key, value)
	}
	for _, value := range courseMap {
		fmt.Println(value)
	}

	//var courseMap2 map[string]string //nil空map
	//courseMap2 := map[string]string{} //map类型想要往里面放值 比如要先初始化
	courseMap2 := make(map[string]string, 3) //make是内置函数  主要用于初始化(slice.map.channel)
	courseMap2["name"] = "Zinon"
	fmt.Println(courseMap2)

	//map必须初始化才能使用 有两种初始化方式
	// 1. make(map[string]string)
	// 2. map[string]string{}
	// slice可以不初始化 但空切片 == nil
	//var courseSlice = map[string]string{}
	//if nil == courseSlice {
	//	fmt.Println("yes you are slice is nil")
	//}
	//courseSlice = append(courseSlice, "啊吧")
	//fmt.Println(courseSlice)
	courseMap["java"] = "java深度学习"
	if _, ok := courseMap["java"]; !ok {
		fmt.Println("not found java")
	} else {
		fmt.Println("isset java")
	}
	//删除元素
	delete(courseMap, "java")
	fmt.Println(courseMap)
	// !! map不是线程安全的
	syncMap := sync.Map{}
	fmt.Println(syncMap)
}
