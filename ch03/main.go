package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//长度计算 go的字符计算分为字节长度/数量 和字符长度 数量
	//如果你想知道一个包含中文的字符串长度要转切片 如果你只有英文 可以通过len直接获取
	name := "Zinon"
	title := "你好 Zinon"
	bytes := []byte(title)
	fmt.Println(len(name))
	fmt.Println(len(bytes))

	//转义符
	//\r \t \n \' \" \? \\
	courseName := "Go\r\n体系课"
	fmt.Println(courseName)
	fmt.Print(courseName + "\r\n")

	username := "Zinon"
	age := 17
	var numbers = []int{15, 16, 17, 18, 19}
	//out := "欢迎您" + username + ", 回到到家"
	fmt.Printf("欢迎您" + username + ", 回到到家 恭喜您已经" + strconv.Itoa(age) + "岁了\r\n") //第一种方式
	fmt.Printf("欢迎回到到家,尊敬的%s 恭喜您已经%d岁了\r\n", username, age)                    //第二种方式
	userMsg := fmt.Sprintf("欢迎回到到家,尊敬的%s 恭喜您已经%d岁了", username, age)            //第三种方式
	fmt.Println(userMsg)
	fmt.Printf("欢迎回到到家,尊敬的%s 恭喜您已经%T岁了\r\n", username, numbers)

	var build strings.Builder //第四种方式
	build.WriteString("欢迎您")
	build.WriteString(username)
	build.WriteString(", 回到到家 恭喜您已经")
	build.WriteString(strconv.Itoa(age))
	build.WriteString("岁了")
	re := build.String()
	fmt.Println(re)
}
