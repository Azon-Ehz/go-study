package main

import (
	"fmt"
	"strconv"
)

func localIncr() func() int { //一个函数中可以访问内部的匿名函数的变量 但一个函数的局部变量不可以给另一个函数访问
	local := 0
	return func() int { //闭包函数可以"记住"并访问其词法作用域中的变量，即使函数在其定义的作用域之外执行。
		local += 1
		return local
	}
}

// 函数参数传递的时候 int float string 是值传递 go语言全部是值传递
func sum(desc string, items ...int) (ret string, error error) {
	sum := 0
	for _, item := range items {
		sum += item
	}
	ret = desc + strconv.Itoa(sum)
	error = nil
	return ret, error
}

func cal(myFunc func(items ...int) int) int {
	return myFunc()
}

func callback(x int, y int, f func(int, int)) {
	f(x, y)
}

func main() {
	//go 函数 支持普通函数 匿名函数 闭包函数
	/*
		go 中函数是一等公民
		1. 函数变身可以当成变量
		2.匿名函数 闭包函数
		3.函数可以满足接口
	*/
	//sum()不带括号 则可以赋值给变量
	funVar := sum
	a := 1
	b := 2
	sum, _ := funVar("总数是:", a, b, 1, 1, 1, 1)
	fmt.Println(sum)

	//函数的入参可以是函数 可以写为闭包函数
	//在函数体内可以调用作为入参的函数
	cal(func(items ...int) int {
		sum := 0
		for _, item := range items {
			sum += item
		}
		return sum
	})

	//如果你调用的函数返回的还是一个函数 需要执行需要这样调用 cal()()
	//闭包函数
	localFunc := func(a int, b int) {
		fmt.Printf("callbackaaa called with a+b= %d \r\n", a+b)
	}
	callback(1, 2, localFunc)
	//闭包的需求 希望有个函数返回的都是每次调用增加之后的值
	nextFunc := localIncr()
	for i := 0; i < 5; i++ {
		fmt.Println(nextFunc())
	}
}
