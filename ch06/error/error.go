package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {
	//1. defer 需要在panic前定义 recover只有在defer调用的函数中才能生效
	//2. recover处理异常后 逻辑不会恢复到panic的那个点去
	//3. 多个defer会形成栈 后定义的defer会先执行
	defer func() {
		//recover的意义就在于捕获程序中的调用的panic 不论是主动还是被动
		// 以提高程序的健壮性 避免出现问题
		if r := recover(); r != nil {
			fmt.Println("Recovered in A", r)
		}
	}()

	//panic()会导致你的程序退出
	//在服务启动过程中 部分依赖组建必须启动时 比如mysql联通、日志文件存在等，在检查依赖时如果出现问题，主动调用panic 去抛出异常退出程序
	//当你的程序出现异常 比如报错等问题 导致程序退出 实际上是被动的调用了panic
	panic(" this is a panic")
	var names map[string]string
	names["go"] = "go开发工程师"
	return 0, errors.New(" this is an error")
}

func main() {
	// error \ panic \ recover
	// go语言错误处理的理念 当函数可能出错时 其他语言是使用try:catch去包住他
	// go 语言是通过函数返回的值 error 判断是否为nil 来告诉调用函数者 是否正常执行成功 也就会导致代码中会有很多 if err != nil
	// go 设计者认为你的error只要返回了 就要处理 也就是防御性编程
	if _, err := A(); err != nil {
		fmt.Println(err)
	}
}
