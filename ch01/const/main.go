package main

import "fmt"

func main() {
	// 常量， 定义的时候就指定的值 不能修改 常量尽量全部大写 多单词通过 _ 连接
	const PI float32 = 3.1415926 //显式定义
	const PI2 = 3.1415926        //隐式定义

	const (
		UNKNOWN = 1
		FEMALE  = 2
		MALE    = 3
	)

	const (
		x int = 1
		y
		s = "abc"
		z
		c
	)
	fmt.Println(UNKNOWN, FEMALE, MALE)
	fmt.Println(x, y, s, z, c)

	/*
		常量类型只能定义Bool/数值(整数，浮点数和复数) 和 字符串
		不曾使用的常量，没有强制使用要求
		显示指定类型的时候，必须确保常量左右定义的类型一致
	*/
}
