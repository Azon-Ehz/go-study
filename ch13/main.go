package main

func main() {
	//go path 和 go module
	/**
	早期的go版本使用go path来管理包
	新版的则使用go module
	判断是使用gopath还是gomod可以通过项目目录下是否存在go.mod来判断
	*** go path在引用包的时候 必须将项目文件夹放在go目录下的src文件夹内 否则会无法引用包
	现在直接创建go文件默认是go.mod的方式	早期的goland会有一个goPath选项
	如果你创建的是goPath文件 则引用包可能会提示 (包名 is not in GOROOT)错误
	*** 那么这个时候你就需要去go env中 检查 GO111MOUDLE 是否是off 如果是on 改为off
	*** go env -w GO111GOMOUDLE=off 因为你现在使用的是goPath模式

	goPath中 import查询顺序
	1.gopath/src 这个目录下是否有包
	2.goroot/src 如果gopath没有 会再去goroot里面找

	goPath项目如何转go module
	1.编辑器设置-启动GO Modules
	2.配置proxy GOPROXY="https://goproxy.cn,direct"
	3.手动新建go.mod文件 package项目名 go 1.19(版本)
	4.开启go env中的 GO111MODULE= on


	包最被吐槽的三个点
	1.包管理
	2.异常处理
	3.泛型

	总结
	1. 能用go module就用go module 不要去考虑以前是开发模式
	2.即是你当前项目是以前的开发模式 也可以设置为现在的go module
	3.vendor补丁模式暂时不提 通常用于用于不同项目不同版本包的问题
	*/
}
