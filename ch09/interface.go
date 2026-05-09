package main

import "fmt"

// 接口的定义
type Duck interface {
	//方法的声明
	Gaga() //返回string
	Walk()
	Swimming()
}

// 接口的实现 需要实现所有的方法
type pskDuck struct {
	legs int
}

func (pd *pskDuck) Gaga() {
	fmt.Println("嘎嘎")
}
func (pd *pskDuck) Walk() {
	fmt.Println("鸭子走路")
}
func (pd *pskDuck) Swimming() {
	fmt.Println("Swimming")
}

func main() {
	/*
		go语言的接口 鸭子类型
		go语言处处都是接口 处处都是鸭子类型 duck typing
		鸭子类型强调的是事物的外部行为 而不是内部的结构
	*/

	var d Duck = &pskDuck{}
	d.Walk()
}
