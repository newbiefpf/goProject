package service

import (
	"github.com/gin-gonic/gin"
	"goProject/models"
)

// UserInfo
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /ping1 [get]
func UserInfo(c *gin.Context) {
	models.GetUser()
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
