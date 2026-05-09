package main

import "fmt"

func main() {
	week := 1

	switch week {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
		break
	case 6, 7:
		fmt.Println("休息日")
		break
	default:
		fmt.Println("未知的")
	}
	/**
	switch有两种写法 也就是是否在switch后跟变量
	这将会决定你是否可以在条件中写表达式
	*/
	switch {
	case week < 5:
		fmt.Println("工作日")
		break
	case week > 5 && week <= 7:
		fmt.Println("休息日")
		break
	default:
		fmt.Println("未知的")
	}
}
