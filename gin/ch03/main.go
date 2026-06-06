package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": context.Param("id"),
		"action":  context.Param("action"),
	})
}
func goodsDetail(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{

		"message": context.Param("id"),
		"action":  context.Param("action")})
}
func goodsAdd(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "goodsAdd",
	})
}
func goodsDelete(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "goodsDelete",
	})
}
func main() {

	gi := gin.Default()
	goodsGroup := gi.Group("/goods")
	{
		goodsGroup.GET("", goodsList)
		goodsGroup.GET("/:id/*action", goodsDetail)
		goodsGroup.POST("", goodsAdd)
		goodsGroup.DELETE(":id", goodsDelete)
	}
	gi.Run(":8383")
}
