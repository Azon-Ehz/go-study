package main

import (
	"fmt"
	"strconv"
)

func main() {
	//同类型不同范围转换
	var a int8 = 12
	var b = uint8(a)

	//浮点与整数转换 会丢值
	var f float32 = 3.14
	var c = int32(f)
	fmt.Println(b, c)

	//浮点同类型转换
	var f64 = float64(a)
	fmt.Println(f64)

	type IT int //类型别名 用于在转换时使用 便于开发
	var abc = IT(a)
	fmt.Println(abc)

	//字符串转Int类型 需要用到strconv包
	var istr = "233"
	var myint int
	myint, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println("convert Atoi err:", err)
		return
	}
	fmt.Println(myint, err)

	var myint2 = 32
	mystr := strconv.Itoa(myint2)
	fmt.Println(mystr)

	//字符串转基础类型
	var myf, err2 = strconv.ParseFloat("3.1415926535", 64)
	if err2 != nil {
		fmt.Println("convert ParseFloat err:", err2)
	}
	fmt.Println(myf)

	ParseInt, err3 := strconv.ParseInt("-42", 10, 64)
	if err3 != nil {
		fmt.Println("convert ParseInt err:", err3)
	}
	fmt.Println(ParseInt)

	ParseBool, err3 := strconv.ParseBool("1")
	if err3 != nil {
		fmt.Println("convert ParseInt err:", err3)
	}
	fmt.Println(ParseBool)

	//基础类型转字符串
	FormatBool := strconv.FormatBool(true)
	fmt.Println(FormatBool)
	fmt.Println(strconv.FormatInt(42, 16))
	//fmt E会携带E+00 f不携带
	fmt.Println(strconv.FormatFloat(3.1415926, 'f', -1, 64))
}
