package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "goProject/docs"
	"goProject/router/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/ping", service.Ping)
	r.GET("/ping1", service.UserInfo)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	return r
}
