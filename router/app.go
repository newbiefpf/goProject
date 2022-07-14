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
	// /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/login", service.Login)
	r.POST("/sendCode", service.SendCode)
	//接口方法
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", service.Ping)

	}
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	return r
}
