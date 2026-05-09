package main

/*
rpc 相关知识概念
*/

type Company struct {
	Name    string
	Address string
}

type Employee struct {
	Name    string
	Phone   string
	Company Company
}

func RpcPrintln(employee Employee) {

	//rpc中的第二个点 传输协议/数据编码协议
	//http1.x(主流) http2.0协议(新推出的 支持长连接)
	//http协议的底层还是tcp http1.0的性能问题主要体现在一次性 一旦结果返回 连接断开
	//所以我们一般要么使用 tcp自行封装编写一个支持长连接的myHttp 要么就使用http2.0
	
	/*
		客户端
			1.建立连接 - 发起请求
			2.发送employee对象序列化成json - 序列化
			3.发送json字符串 - 调用成功后实际上你接收到的是一个二进制的数据(经序列化)
			4.等待服务器发送结果
			5.将服务器返回的数据解析成可用的对象 - 反序列化
		服务端
			1.监听服务端口
			2.读取客户端发送的数据 -二进制的json
			3.对数据进行解析 - 反序列化
			4.开始业务逻辑处理
			5.将处理完成的结果序列化成json-二进制数据-序列化处理
			6.将数据返回
		序列化、反序列化是可以选择的 不一定采用json 还有 xml/protobuf/msgpack
	*/
}

func main() {

}
