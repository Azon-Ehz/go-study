package main

import "fmt"

type Person struct {
	name string
	age  int

	address struct {
		city string
	}
}

type Student struct {
	//结构体的嵌套
	//p Person

	//匿名嵌套
	Person
	//两种嵌套方式的区别在于取值的时不需要带第一种方式的 s.p.name 而是直接 s.name
	name  string
	score float32
}

// func(接收的结构体) funcName(如参) (返回){函数体}
// 接收器有两种形态 * 有无指针表示标识
// 如果你的函数想要修改结构体的值 但结构体的值很大 就不建议使用值传递 而使用指针传递
func (s *Student) print() {
	s.age = 333
	s.name = "zhang"
	fmt.Printf("name: %s, age: %d\n", s.name, s.age)
}
func (s Student) print2() {
	fmt.Printf("name: %s, age: %d\n", s.name, s.age)
}

func main() {

	s := &Student{
		Person{
			"Zinon",
			18,
			struct{ city string }{city: "北京"},
		},
		"Zinon3",
		99.2,
	}
	s.print()
	s.print2()
	//
	//s.name = "Zinon2"
	//fmt.Println(s.name)
	//s.address.city = "上海"
	//fmt.Println(s.address.city)
	//s.name = "ZinonZ"
	//s.age = 32
	//s.print()
}
