package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// go modules依赖管理添加和删除依赖项
// go list -m -versions github.com/gin-gonic/gin 该依赖的可用版本
// 使用 go get  github.com/gin-gonic/gin@v1.12 可以获取指定版本 并覆盖当前本地版本
// go list -m all 当前项目中所依赖项有哪些
// go get -u 升级到最新的次要版本或修订版本
// go get -u=patch 升级到最新的修订版本
// get get client_proxy@version go get 会修改 go mod 文件

// A项目中的 github仓库是project-A go.mod中的定义是 github.com/bobby/A
// B项目由于依赖A项目 所以需要下载A项目的依赖 import的地址应该是github.com/bobby/project-A. go get的时候就会出现go.mod和项目路径不一样的问题
// 这个时候可以使用replace 命令具体是 go mod edit -replace github.com/bobby/A=github.com/bobby/project-A@v1.0.0
// 那么当代码执行完成后 go.mod 文件中就会新增一行 replace github.com/bobby/A => github.com/bobby/project-A v1.0.0
// go get/mod/install
func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "222",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
