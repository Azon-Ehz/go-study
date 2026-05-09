package main

import "fmt"

/*
goto  语句可以让我的代码跳到指定的代码块中运行 直接跳出原有代码  很灵活 但因此很少用

go语言的gto语句可以实现程序的条件 goto最多的使用场景是在程序的错误处理 异常处理
也就是当程序出现错误的时 立刻跳转到指定的标签代码块中统一处理
*/
func main() {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				goto gover
			}
			fmt.Println(i, j)
		}
	}
	//定义代码块别名
gover:
	{
		fmt.Println("goto初体验")
	}
}
