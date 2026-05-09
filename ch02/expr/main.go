package main

import (
	"fmt"
	"strconv"
)

func main() {
	//算数运算符 + - * / % ++ -- += -=
	a, b := 1, 2
	astr, bstr, cstr := "hello", " ", "zinon"
	c, d := 3, 10
	fmt.Println(a + b)
	fmt.Println(astr + bstr + cstr)
	fmt.Println(d % c)
	a++
	a += 1
	a = a + 1
	a--
	a -= 1
	a = a - 1
	fmt.Println(a)

	// 逻辑运算符 && || !
	abool := true
	bbool := false
	if !abool && !bbool {
		fmt.Println("Success")
	} else {
		fmt.Println("fail")
	}

	//位运算符 将数值转二进制后运算每位
	A, B := 60, 13
	//0011 1100
	//0000 1101
	fmt.Println(strconv.FormatInt(int64(A), 2))
	fmt.Println(strconv.FormatInt(int64(B), 2))
	fmt.Println(A & B) //按位与 同等于逻辑且
	fmt.Println(A | B) //按位或 同等于逻辑或
	fmt.Println(A ^ B) //按位异或 ⚠️只有当对比的两位不同时 才为真

	//移位运算符 暂时看不懂
	fmt.Println(A << B)
	fmt.Println(A >> B)

	//地址符 &返回变量在内存中的地址 *
	e := &A    //取地址
	var f *int //指针
	fmt.Println(e)
	fmt.Println(f)
}
