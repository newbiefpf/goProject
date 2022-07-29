package controller

import "github.com/gin-gonic/gin"

// Ping
// @Tags A心跳
// @Summary 测试联通
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "pong",
	})
}
