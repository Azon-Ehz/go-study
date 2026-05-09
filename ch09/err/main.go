package main

import "fmt"

func mPrint(data ...interface{}) {
	for _, v := range data {
		fmt.Println(v)
	}
}

type myInfo struct{}

func (mi *myInfo) Error() string {
	return "我不是Error"
}

func main() {
	var data = []interface{}{
		"bobby", "Zinon", "hello", "Golang", "gin", 8832, 22.3,
	}
	mPrint(data...)
	var data2 = []string{
		"bobby", "Zinon", "hello", "Golang", "gin",
	}
	var data3 []interface{}
	for _, value := range data2 {
		data3 = append(data3, value)
	}
	mPrint(data3...)

	err := myInfo{}
	fmt.Println(err.Error())
}
