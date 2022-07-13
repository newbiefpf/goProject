package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goProject/conf"
	"goProject/controller/user"
	"log"
	"net/http"
)

// NoResponse 请求的URL不存在，返回404
func NoResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error":  "404, page not exists!",
		"status": 404,
	})
}

//Route 请求的url路由
func Route() {
	//路由分组
	router := gin.Default()
	router.Use(gin.Recovery())
	port := conf.BlogConfig.Server.Port
	log.Printf("服务端口: %d ", port)

	v1 := router.Group("/v1")
	{
		//调用方法
		v1.GET("/get", user.FindUserInfo)
	}
	// 路径映射
	log.Printf("服务端口: %d ", port)
	router.Run(":" + fmt.Sprintf("%d", port))
}
