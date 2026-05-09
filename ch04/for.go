package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
	for 初始值;条件;步长  {
	  逻辑
	}
*/
func main() {

	sum := 0
	var i = 1
	for i <= 100 {
		fmt.Println(i)
		sum += i
		i++
	}
	fmt.Println(sum)

	//99乘法
	for k := 1; k <= 9; k++ {
		for key := 1; key <= k; key++ {
			var build strings.Builder
			build.WriteString(strconv.Itoa(key))
			build.WriteString("*")
			build.WriteString(strconv.Itoa(k))
			build.WriteString("=")
			build.WriteString(strconv.Itoa(k * key))
			re := build.String()
			fmt.Print(re + " ")
		}
		fmt.Println()
	}
	//遍历
	// 字符串的索引 	字符串对应的索引字符值的拷贝(value） 如果不写key 返回的是索引
	// 数组的索引 	索引对应值的拷贝					如果不写key 返回的是索引
	// 切片的索引 	索引对应值的拷贝					如果不写key 返回的是索引
	// map的索引 	value是key对应值的拷贝			如果不写key 返回的是map的值
	// channel 		value返回的是 channel接收的数据
	name := "Ziono Like Go 哈哈"
	for _, index := range name {
		fmt.Printf("%c", index)
	}
	nameRune := []rune(name)
	fmt.Println("\r\n" + strconv.Itoa(len(nameRune))) //16个字节

	for init := range nameRune {
		//fmt.Printf("%c", name[init]) // 直接用下标会有问题 因为中文占3个字节 索引获取值异常 会取中文的内部字节 而不是中文本身 自然是乱码
		fmt.Printf("%c \r\n", nameRune[init])
	}
	for ik := 0; ik < len(nameRune); ik++ {
		fmt.Printf("%c \r\n", nameRune[ik])
		//break 直接退出当层循环
		//continue // 跳过当次循环
	}

	for l := 0; l <= 10; l++ {
		if l == 5 {
			continue
		}
		if l > 9 {
			break
		}
		fmt.Println(l)
		//休眠
		time.Sleep(1 * time.Second)

	}
}
