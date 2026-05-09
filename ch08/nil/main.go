package main

import "fmt"

type Person struct { //复杂的结合体
	name string
	age  int
	f    *int
}

func main() {
	/*
		在go语言中不同类型的空值不一样
		bool false
		numbers 0
		string ""
		pointer nil
		slice nil
		map nil
		channel interface function nil

		struct 默认值不是nil 默认值是具体字段的默认值
	*/

	p1 := Person{
		name: "Zinon",
		age:  20,
	}
	p2 := Person{
		name: "Zinon",
		age:  19,
	}
	if p1 == p2 { //struct在对比时 必须是每个字段都相同 结果才会true
		fmt.Println("p1 equals p2")
	} else {
		fmt.Println("p1 != p2")
	}

	//var ps []Person             // nil slice
	//var ps2 = make([]Person, 0) //empty slice
	//if ps == nil {
	//	fmt.Println("y")
	//} else {
	//	fmt.Println("n")
	//}
	//var m map[string]string             // nil map
	var m2 = make(map[string]string, 0) //empty map
	if m2 == nil {
		fmt.Println("y")
	} else {
		fmt.Println("n")
	}

}
