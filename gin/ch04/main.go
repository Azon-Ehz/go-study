package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/:id/:name", func(context *gin.Context) {
		var person Person
		if err := context.ShouldBindUri(&person); err != nil {
			context.Status(444)
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"id":   person.ID,
			"name": person.Name,
		})

	})
	router.Run(":8383")
}
