package main

import (
	"fmt"
	"strings"
)

func add(a, b interface{}) any {
	switch a.(type) {
	case int:
		ai, ok := a.(int)
		if !ok {
			panic("a is not int type")
		}

		bi, ok := b.(int)
		if !ok {
			panic("b is not int type")
		}
		return ai + bi
	case int32:
		ai, ok := a.(int)
		if !ok {
			panic("a is not int type")
		}

		bi, ok := b.(int)
		if !ok {
			panic("b is not int type")
		}
		return ai + bi
	case float64:
		ai, ok := a.(float64)
		if !ok {
			panic("a is not int type")
		}

		bi, ok := b.(float64)
		if !ok {
			panic("b is not int type")
		}
		return ai + bi
	case string:
		ai, ok := a.(string)
		if !ok {
			panic("a is not int type")
		}

		bi, ok := b.(string)
		if !ok {
			panic("b is not int type")
		}
		return ai + bi
	default:
		panic("暂不支持其它类型")
	}

}

func main() {
	var a = "Hello"
	b := " Zinon"
	str := add(a, b)
	str2, _ := str.(string)
	fmt.Println(strings.Split(str2, " "))
	fmt.Println(str2)
}
