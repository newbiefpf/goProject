package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "goProject/docs"
	"goProject/router/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	// /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//v1 := r.Group("/v1")
	//{
	//	v1.GET("/ping", controller.Ping)
	//
	//}
	r.GET("/ping", controller.Ping)
	//r.POST("/login", controller.Login)
	r.POST("/sendCode", controller.SendCode)
	//接口方法

	r.Run(":8888") // 监听并在 0.0.0.0:8888 上启动服务
	return r
}
