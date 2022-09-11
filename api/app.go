package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	articleApi "goProject/api/article"
	userApi "goProject/api/user"
	"goProject/middleware"
	"goProject/middleware/logMiddleware"
)

func Init() {
	r := gin.New()
	r.Use(logMiddleware.LoggerToFile())
	//登录
	r.POST("/login", userApi.Login)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/v1").Use(middleware.JWTAuthMiddleware())
	{

		v1.POST("/connection", userApi.Pong)

		//发布文章
		v1.PUT("/publish", articleApi.CreatArticle)
		//创建用户
		v1.PUT("/creatUser", userApi.CreatUser)

		v1.GET("/connection", userApi.Ping)
		//获取文章
		v1.GET("/articleList", articleApi.GetArticleList)

	}
	r.Run(":8888")
}
