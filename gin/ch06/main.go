package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"awesomeProject/gin/ch06/proto"

	googleProto "google.golang.org/protobuf/proto"
)

func main() {
	router := gin.Default()
	router.POST("/post", testPost)
	router.POST("/proto", returnProtobuf)
	router.GET("/purejson", returnPureJson)
	router.Run(":8383")
}

func testPost(c *gin.Context) {
	var msg struct {
		Name    string `json:"username"`
		Message string
		Number  int
	}
	msg.Name = "Zinon"
	msg.Message = "这个一条测试Json-struct"
	msg.Number = 20

	c.JSON(http.StatusOK, msg)
}
func returnProtobuf(c *gin.Context) {
	course := []string{"python", "PHP", "Golang", "微服务"}
	user := &proto.Teacher{
		Name:   "Zinon",
		Course: course,
	}
	data, err := googleProto.Marshal(user)
	if err != nil {
		panic(err)
	}
	err = googleProto.Unmarshal(data, user)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, user)

}
func returnPureJson(c *gin.Context) {
	c.PureJSON(http.StatusOK, gin.H{
		"html": "<b>Hello World</b>",
	})
}
