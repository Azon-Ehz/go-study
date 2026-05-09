package main

import "fmt"

// 全局变量
var (
	name = "Zinon"
	age  = 25
	sex  = true
)

func main() {
	//go是静态语言和动态语言的差异很大
	// 1.变量必须先定义后使用 2.变量定义必须有类型  3.类型定下来后不能改变
	//局部变量´
	// var name int
	// name = 1
	//var age = 1
	//age := 1
	// go语言中 局部变量定义了一定要使用
	//fmt.Println(age)

	//多变量定义
	var user1, user2, user3, user4 = 1, "2", '3', '4'
	fmt.Println(user1, user2, user3, user4)

	/*
		变量必须定义才能使用
		go语言是静态语言 要求变量的值要和赋值的类型一致
		变量名 同一个代码块(局部) 不能重复
		简洁变量不能用于全局
		变量是有零值的 var (age int, name string) 0和空字符 默认值
		局部变量定义了不使用会报错 全局不会
	*/
}
