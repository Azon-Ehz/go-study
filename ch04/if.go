package main

import "fmt"

/*
	if 布尔表达式 {
	  逻辑
	}
*/
func main() {

	age := 18

	if age < 18 {
		fmt.Println("未成年")
	} else if age == 18 {
		fmt.Println("刚好成年")
	} else {
		fmt.Println("成年")
	}
	/**
	上下两种if上的区别就在于
	多一次比较数据 上面的判断逻辑是直接else 不比较 而下面的则是需要比对 性能相比差一点
	*/
	if age < 18 {
		fmt.Println("未成年")
	}
	if age == 18 {
		fmt.Println("刚好成年")
	}
	if age > 18 {
		fmt.Println("成年")
	}
}
