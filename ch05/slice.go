package main

import "fmt"

func main() {
	//动态语言和静态语言 数组的定义差别在长度固定 静态语言是不允许在定义后对数组长度做调整的 比如append pop等操作

	//go 是折中的 数组长度不一样 数组的类型都不一样 实际上是在弱化数组的功能 也就是降低数组的使用率
	//而slice(切片)就是go折中方案的体现 但要记住的底层还是数组
	var courses []string
	fmt.Printf("%T \r\n", courses) //[]string 切片

	courses = append(courses, "Go")
	courses = append(courses, "Grpc")
	courses = append(courses, "Gin")
	fmt.Println(courses)
	fmt.Println(courses[2])

	//切片的初始化方式有四种 1.从数组直接创建 2. 从切片创建切片 3.var courses33 = []string{} 4. make
	//1. 从数组直接创建
	allCourses := [5]string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	allCoursesLen := len(allCourses)
	allCoursesSlice := allCourses[0:allCoursesLen] //左闭右开【)
	fmt.Println(allCoursesSlice)

	//2.从切片创建切片
	allCoursesSlice2 := allCourses[0:len(allCourses)] //左闭右开【)
	fmt.Println(allCoursesSlice2)

	//3. 使用[]string{}
	allCoursesSlice3 := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	fmt.Println(allCoursesSlice3)

	//4. 使用make
	allCoursesSlice4 := make([]string, allCoursesLen)
	allCoursesSlice5 := make([]string, 3)
	var allCoursesSlice6 []string
	fmt.Println(len(allCoursesSlice4))
	fmt.Println(len(allCoursesSlice5))
	allCoursesSlice6 = append(allCoursesSlice6, "GoLand")
	fmt.Println(allCourses[:]) //llCourses[start:end] 两个都是选填 都不填则是全部

	slices := []string{"mysql", "elasticsearch"}
	slices = append(slices, "Zinon", "study")

	mySlice := append(allCourses[:2], allCourses[3:]...)
	fmt.Printf("%T \r\n", mySlice)
	mySlice = mySlice[:3] //切片长度可变
	fmt.Printf("%T \r\n", mySlice)
	fmt.Println(mySlice)

	//复制slice 切片
	//mySlice2 := mySlice
	mySliceCopy1 := mySlice[:]
	//fmt.Println(mySlice2, mySlice3)
	//var mySlice4 []string 没有长度的切片如果进行copy 是不会赋值的 不会扩容长度
	var mySliceCopy = make([]string, len(mySliceCopy1))
	copy(mySliceCopy, mySlice)
	fmt.Println("-----------------------------------------------")

	fmt.Println(mySliceCopy, mySliceCopy1)
	fmt.Println("-----------------------------------------------")
	mySlice[0] = "Java"
	fmt.Println(mySliceCopy1) // := mySlice[:] 这种方式 如果你修改了源数据 那么被赋值的数据也会跟着修改 无关代码执行顺序
	fmt.Println(mySliceCopy)  // copy函数 即使源数据被修改 copy后的值也不会被影响
}
