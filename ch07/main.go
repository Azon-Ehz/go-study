package main

import (
	"fmt"
	"strconv"
)

type myInt int //在编译过程中 别名将会直接被替换成相应的数据类型

func (mi myInt) string() string {
	return strconv.Itoa(int(mi))
}

func main() {
	//type 关键字
	/*
		1. 定义结构体
		2. 定义接口
		3. 定义类型别名
		4. 类型定义
		5. 类型判断
	*/
	//别名实际上是为了更好的理解代码
	var i myInt = 12
	fmt.Println(i.string() + "aaaa")
	fmt.Printf("%T\n", i)
	var a interface{} = "abc"
	switch a.(type) { //获取真正的类型
	case string:
		fmt.Println("string is a string")
		break
	case int:
		fmt.Println("int is a int")
		break
	}
	m := a.(string) //转换类型
	fmt.Printf("%T\n", m)
}
