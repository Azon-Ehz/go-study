package main

import "fmt"

// 人员信息存储
//var persons [][]string 二维数组
//var persons [][]interface{} //interface

// 类型集合
type Person struct {
	name    string
	age     int
	address string
	height  float64
}

func main() {
	//persons = append(persons, []string{"Zinon", "18", "golang", "1.73"})
	//persons = append(persons, []interface{}{"Zinon", 18, "golang", 1.73}) //断言

	//初始化结构体
	person := Person{}
	person.name = "Jack"
	person.age = 18
	person.address = "Jack"
	person.height = 1.73
	fmt.Println(person)

	var persons6 Person
	persons6.name = "Jack"
	persons6.age = 18
	persons6.address = "Jack"
	persons6.height = 1.9
	fmt.Println(persons6.name)

	person2 := Person{"Zinon", 18, "Shanghai", 1.74} //必须全部属性赋值初始化
	fmt.Println(person2)

	person3 := Person{ // 可以部分属性赋值
		name:    "Jack",
		age:     18,
		address: "Jack",
		height:  1.93,
	}
	fmt.Println(person3)

	var persons4 []Person
	persons4 = append(persons4, person)
	persons4 = append(persons4, person2)
	persons4 = append(persons4, person3)
	persons4 = append(persons4, Person{ // 可以部分属性赋值
		name:    "Jack",
		age:     18,
		address: "Jack",
		height:  1.93,
	})
	fmt.Println(persons4)

	persons5 := []Person{
		{"Zinon", 18, "Shanghai", 1.74},
		{
			age:     18,
			address: "Jack",
		},
		{
			name:    "Jack",
			age:     18,
			address: "Jack",
		},
		{
			name:   "Jack",
			age:    18,
			height: 1.93,
		},
		{
			address: "Jack",
			height:  1.93,
		},
	}
	fmt.Println(persons5)

	//匿名结构体，匿名函数
	address := struct {
		//定义结构体
		province string
		city     string
		address  string
		num      int
	}{ //实例化
		province: "上海市",
		city:     "静安区",
		address:  "静安路199号",
	}
	fmt.Println(address.num)
}
