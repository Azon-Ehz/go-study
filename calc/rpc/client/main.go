package main

import (
	"encoding/json"
	"fmt"

	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int "json:data"
}

func Add(a, b int) int {

	req := HttpRequest.NewRequest()
	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b))
	body, _ := res.Body()
	//fmt.Println(string(body))
	RspData := ResponseData{}
	_ = json.Unmarshal(body, &RspData)
	return RspData.Data
}

func main() {
	fmt.Println(Add(2, 3))
}
