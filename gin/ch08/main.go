package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "hello2222")
		c.Next()
		end := time.Since(t)
		fmt.Printf("耗时 %V\n", end)
		status := c.Writer.Status()
		fmt.Println("状态", status)
	}
}

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			fmt.Println(2, 2, k, v)

			if k == "x-token" {
				token = v[0]
			}

		}
		if token != "Zinon" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"massage": "未登录",
			})
			c.Abort()
		}
		c.Next()
	}
}

func main() {
	router := gin.New()
	//使用logger、Recovery中间件中间件 ——————全局使用
	router.Use(MyLogger(), TokenRequired())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8383")
}
