package course

/*
client_proxy 用来组织源码 是多个go源码的集合 是代码复用的基础 fmt, os, io
每个源码文件开始编写前都需要声明所属的package
java, go 通过package声明 而php通过namespaces声明 python通过文件名隐性的被动声明
go在定义时 同一文件夹下可以有多个源码文件 但必须都属于同一个包(client_proxy)
*/
func GetCourse(c Course) string {
	return c.Name
}
