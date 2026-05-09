package main

import (
	"fmt"
)

func a() (int, bool) {
	return 0, false
}

var name = "Zinon" //全局变量

func main() {
	//匿名变量 , 变量作用域
	var _ int
	_, ok := a()
	if ok {
	}

	var mainName = "main"
	fmt.Println(mainName)

	//iota 特殊常量-可以被编译器修改的常量
	const ( //在批量定义 赋值iota时 iota内部存在计数器 会自行提供默认值 从0开始递增
		ERROR0 = iota
		ERROR1
		ERROR2
		ERROR3

		ERROR4 = "h4" //iota内部仍然计数
		ERROR5

		ERROR6 = iota //iota直接计数批量定义 而非定义iota的数量
	)
	const a = iota //
	fmt.Println(ERROR1, ERROR2, ERROR3, ERROR6)

	/**
	在批量定义 赋值iota时 iota内部存在计数器 会自行提供默认值 从0开始递增
	如果iota中主动定义了其他常量值 iota内部仍然计数 这里的计数器 直接计数当前批量定义的常量数量 而非主动定义iota的数量
	每次出现const(重新发起定义) iota归零
	常量批量定义iota过程中如果中断了 那么定义必须显式的恢复
	iota的自增值(默认int) 但实际是计数器的值 你定义的整数/字符串 不会影响iota的值
	iota可以简化const的定义
	*/
	{
		localName := "local"
		fmt.Println(localName)
	}

	if name == "Zinon" {
		var mname string
		//这个变量的作用域只有这个 if内
		mname = "imooc"
		fmt.Println(mname)
	}
	fmt.Println(name)
}
