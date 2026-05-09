package main

import "fmt"

type Person struct {
	name string
}
type PersonName struct {
	name string
}

func changeName(p *Person) {
	p.name = "imooc"
}

func (p *Person) sayHello() {

}
func (pn *PersonName) sayHello() {

}

func swap(a, b *int) {
	fmt.Println("swap", a, b) //0x14000018088 0x14000018090 1 2
	a, b = b, a               //要使用指针关键符 * 去交换地址指向的值

	fmt.Println("swap", a, b) //0x14000018090 0x14000018088  2 1
}

func main() {
	//希望在结构体传值的时候 在函数中的修改可以体现在变量中
	//p := Person{
	//	name: "Zinon",
	//}
	//changeName(&p)
	//var p1 *Person = &p
	//
	//var po *Person
	//po = &p
	//pi := &p
	//fmt.Println(&po)
	//fmt.Println(&pi)
	//fmt.Printf("%p", &po)
	////指针 与 指针变量 和 值 的概念关系
	///*
	//	如上述代码 指针变量在定义时需要使用*表明
	//	在想获取变量值的时候 需要加上&地址符去获取
	//	如果直接打印指针变量 则会打印出这个变量在内存中的地址如0x14000010240
	//*/
	//p2 := &Person{
	//	"Zinon2",
	//}
	//fmt.Println("\r\n" + p2.name)
	//fmt.Println((*p2).name)
	//与其他语音

	//var a int = 10
	//b := &a

	//var b *int
	//fmt.Println(b)
	//ps := &Person{} //指针初始化第一方法
	//var emptyPerson Person
	//p4 := &emptyPerson //指针初始化第二方法
	//p3 := new(Person)
	//每次初始化的指针地址都不同
	//map channel slice初始化通过make map必须初始化
	//指针初始化通过new
	//指针需要初始化 否则报错 nil pointer
	//fmt.Println(&ps, &p4, &p3)

	c, d := 1, 2
	fmt.Println(&c, &d) //取地址 0x14000018088 0x14000018090
	swap(&c, &d)        //go语言是值传递 这里实际上传递的是变量在内存中的地址
	fmt.Println(c, d)
}
