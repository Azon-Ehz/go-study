package main

import (
	protoHello "awesomeProject/helloworld/proto"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
)

type Hello struct {
	Name    string   `json: "name"`
	Age     int32    `json: "age"`
	Courses []string `json: "courses"`
}

func main() {

	req := protoHello.HelloRequest{
		Name:    "Zinon",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	jsonStruct := Hello{Name: "Zinon", Age: 18, Courses: []string{"go", "gin", "微服务"}}
	jsonRsp, _ := json.Marshal(jsonStruct)
	protoRsp, _ := proto.Marshal(&req) //具体的编码是如何实现的 很多公司没有直接使用protobuf 可以自己尝试开创编码格式
	fmt.Println(len(protoRsp))         //proto 压缩比更高 切片len 9
	fmt.Println(len(jsonRsp))          //json 压缩比更低 切片len 16
	newReq := protoHello.HelloRequest{}
	_ = proto.Unmarshal(protoRsp, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)

	//grpc选择protobuf就是因为他的压缩比性能更高
	//打印的结果虽然不是我们肉眼可以看的，但可以通过protobuf解码获取
	// json 与 protobuf 压缩比接近一倍
	//接下来我们继续看 关于 服务器的存根sub 接口 grpc可以直接生
}
