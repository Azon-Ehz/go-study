package main

import (
	_ "awesomeProject/ch10/init" //自动调用
	//uc "awesomeProject/ch10/user" //包别名 路径
	. "awesomeProject/ch10/user" //将包的内容直接导入当前 main
	"fmt"
)

func main() {
	c := Course{
		Name: "Zinon",
	}
	fmt.Println(GetCourse(c))
	//c := uc.Course{
	//	Name: "Zinon",
	//}
	//fmt.Println(uc.GetCourse(c))

}
