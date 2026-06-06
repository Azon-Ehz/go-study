package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func welcome(c *gin.Context) {
	FirstName := c.DefaultQuery("first", "Zinon")
	LastName := c.DefaultQuery("last", "Zinon")
	c.JSON(http.StatusOK, gin.H{
		"first_name": FirstName,
		"last_name":  LastName,
	})
}

func welcomePost(c *gin.Context) {
	get := c.DefaultQuery("get", "getDefault")
	nickName := c.DefaultPostForm("nick", "Zinon")
	c.JSON(http.StatusOK, gin.H{
		"nick_name": nickName,
		"get":       get,
	})
}
func main() {
	router := gin.Default()
	router.GET("/welcome", welcome)

	router.POST("/post", welcomePost)
	router.Run(":8383")
}
