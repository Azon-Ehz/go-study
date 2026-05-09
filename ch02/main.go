package main

import (
	"fmt"
)

func main() {

	var age1 int8
	var age2 int16
	var age3 int32
	var age4 int64
	var uage5 uint8
	var uage6 uint16
	var uage7 uint32
	var uage8 uint64
	var age int //动态类型 基于操作系统 64位就是int64 32位就是int32  上面的类型很多 但是用起来很麻烦
	fmt.Println(age1, age2, age3, age4, uage5, uage6, uage7, uage8, age)
	//不同范围的数据类型不可以直接转换必须经过
	age1 = int8(age2) //显式转换 但如果age2的值超出int8的范围 会报错

	var f1 float32 //3.4e38
	var f2 float64 //1.8e308
	f1 = 3
	f2 = 3.14
	fmt.Println(f1, f2)

	var b1 byte // byte是uint8的别名 主要用于存储字符 替代别的字符的char
	b1 = 'b'    // 单引号用于char(记住 底层是数字 也就是asccll值 ) 而双引号用于string
	b2 := 98
	fmt.Println(b1)
	fmt.Printf("b1=%c \n", b1)
	fmt.Printf("b2=%c \n", b2)

	var c1 rune // 同byte一样 属于别名 是int32的别名
	c1 = '章'
	fmt.Println(c1)
	fmt.Printf("c1=%c \n", c1)

	var name string
	name = "I am Zinonv "
	fmt.Println(name)
}
