package main

import (
	"fmt"
)

func main() {
	// go 语言的复杂数据类型 数组 切片(slice) map list
	//数组 定义 var name [count]int
	//数组 长度固定 性能高 不同的数组长度意味着不同的数组类型 如果不定义数组长度 则为切片
	var courses [3]string  //courses 只有三个数组元素的字符串数组
	var courses2 [4]string //courses 只有四个数组元素的字符串数组
	courses[0] = "go"
	courses[1] = "grpc"
	courses[2] = "Gin"
	fmt.Println(courses)
	fmt.Printf("%T \n", courses2)

	for index, course := range courses {
		fmt.Println(index, course)
	}

	//定义方式的简写 批量赋值
	courses3 := [3]string{"go", "grpc", "Gin"}
	for index, value := range courses3 {
		fmt.Println(index, value)
	}

	courses2[2] = "Gin"
	fmt.Println(courses2)

	//让编译器根据初始化时提供的元素个数自动推断数组的长度 性能与上面一样
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for _, value := range arr {
		fmt.Println(value)
	}

	arr2 := [...]int{2, 1, 3, 4, 5, 6, 7, 8, 9} // 省略号定义时 必须初始化值
	//数组比较 不同长度的数组不可以直接比较
	//比较时 必须下标 值 全部相同 结果才是真
	if arr == arr2 {
		fmt.Println("x")
	} else {
		fmt.Println("o")
	}

	//多维数组的定义
	var arr3 [3][4]string
	arr3[0] = [4]string{"go", "1h", "Zinon", "go体系课"}    //多维数组赋值
	arr3[1] = [4]string{"grpc", "2h", "Zinon", "grpc入门"} //多维数组赋值
	arr3[2] = [4]string{"go", "1h", "Zinon", "gin高级开发"}  //多维数组赋值
	fmt.Println(arr3)
	//两种遍历方式
	for _, value := range arr3 {
		for _, s := range value {
			fmt.Print(s + " ")
		}
		fmt.Println()
	}
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Print(arr3[i][j] + " ")
		}
		fmt.Println()
	}

	for _, value := range arr3 {
		fmt.Println(value)
	}
}
