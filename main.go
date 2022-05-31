package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func main() {
	r := gin.Default()
	//解析今静态文件
	r.Static("/xxx", "./static")
	//解析模板
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", getIndex)
	r.Run(":9090")
}
