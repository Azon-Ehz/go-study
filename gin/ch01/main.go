package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
	c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, World!",
	})
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello, World!",
	})
}

func main() {
	//实例化一个server对象
	r := gin.Default()
	r.GET("/ping", pong)
	r.Run(":8383")

}
