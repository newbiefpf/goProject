package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	userApi "goProject/api/user"
	"goProject/middleware"
	"goProject/middleware/logMiddleware"
)

func Init() {
	r := gin.New()
	r.Use(logMiddleware.LoggerToFile())
	r.POST("/login", userApi.Login)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/v1").Use(middleware.JWTAuthMiddleware())
	{
		v1.POST("/connection", userApi.Pong)
		v1.PUT("/creatUser", userApi.CreatUser)
		v1.GET("/connection", userApi.Ping)
	}
	r.Run(":8888")
}
