package user

import (
	"github.com/gin-gonic/gin"
	"goProject/model/dbModel"
	"goProject/service/aboutUser"
	"goProject/utility/ginResult"
)

// PONG 心跳
// @Summary 心跳信息
// @Description 心跳信息
// @Tags 心跳接口
// @Accept application/json
// @Produce application/json
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /v1/connection [POST]
func Pong(c *gin.Context) {
	c.JSON(200, ginResult.OK.WithData("PONG!PONG!PONG!"))
}

// PING 心跳
// @Summary 心跳信息
// @Description 心跳信息
// @Tags 心跳接口
// @Accept application/json
// @Produce application/json
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /v1/connection [get]
func Ping(c *gin.Context) {
	c.JSON(200, ginResult.OK.WithData("PING!PING!PING!"))
}

// FindUserByName 创建用户
// @Summary 创建用户信息
// @Description 创建用户
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Router /v1/creatUser [PUT]
func CreatUser(c *gin.Context) {
	var user dbModel.User
	_ = c.BindJSON(&user)
	aboutUser.AddUser(user)
	c.JSON(200, gin.H{
		"message": "测试",
		"data":    user,
	})
}
