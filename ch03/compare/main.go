package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//字符串的比较
	a := "hello"
	b := "bello"
	if a == b {
		fmt.Println("OK")
	}
	if a != b {
		fmt.Println("!=")
	}

	//比较过程（从左到右逐字节） 不是 ASCII 比较
	if a > b {
		fmt.Println("amax")
	}
	if a < b {
		fmt.Println("bmax")
	}

	name := "Zinon-正在努力成为golang开发golang工程师"
	//是否包含
	fmt.Println(strings.Contains(name, "go"))
	fmt.Println(utf8.RuneCountInString(name))

	//指定字符在字符串中出现了多少次
	fmt.Println(strings.Count(name, "n"))

	//字符串转切片 分割字符串
	fmt.Println(strings.Split(name, "-"))

	//查询字符串是否包含前后缀 大小写敏感
	fmt.Println(strings.HasPrefix(name, "Zi"))
	fmt.Println(strings.HasSuffix(name, "工程师"))

	//查询子串在字符串中出现的位置 从0 从左往右开始计数
	fmt.Println(strings.IndexRune(name, []rune(name)[8]))

	//替换目标字符串 可以通过参数n指定全部替换还是部分替换(多个被替换字符情况) -1所有 1/2/3循环次数
	fmt.Println(strings.Replace(name, "golang", "go", -1))

	//大小写
	c := "#ABDCssaaddDDWWWC$"
	fmt.Println(strings.ToUpper(c))
	fmt.Println(strings.ToLower(c))

	//去特殊字符 可批量 可指定
	fmt.Println(strings.Trim(c, "C$%#"))
}
