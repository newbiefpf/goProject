package service

import "github.com/gin-gonic/gin"

// Ping
// @Tags 心跳
// @Summary 测试联通
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
