package controller

import (
	"github.com/gin-gonic/gin"
	rPa "goProject/router/interfaceDataModel"
	"goProject/utility"
	"net/http"
)

// Login
// @Tags 接口实现
// @Summary 登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /login [post]
//func Login(c *gin.Context) {
//
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//	if username == "" || password == "" {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    -100,
//			"message": "必填项为空",
//		})
//	}
//	password = utility.GetMd5(password)
//	data := new(models.Userinfo)
//	err := models.DB.Where("username = ? AND password = ?", username, password).First(&data).Error
//	if err != nil {
//		if err == gorm.ErrRecordNotFound {
//			c.JSON(http.StatusOK, gin.H{
//				"code":    -101,
//				"message": "用户名或密码错误",
//			})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{
//			"code":    -501,
//			"message": "内部错误" + err.Error(),
//		})
//		return
//	}
//	token, err := utility.GenerateToken(string(data.ID), data.Username)
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code":    201,
//			"message": "token 错误" + err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"data": map[string]interface{}{
//			"token": token,
//		},
//	})
//}

// SendCode
// @Tags 接口实现
// @Summary 发送验证码21312
// @Accept application/json
// @Produce application/json
// @Param email body string true "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /sendCode [post]
func SendCode(c *gin.Context) {
	data := rPa.EmailCode{}
	if c.ShouldBind(&data) != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -100,
			"message": "必填项为空",
		})
		return
	}
	//生成唯一的验证码10分钟有效
	code := utility.GetRand()
	//保存验证码
	//err := models.DB.Where("username = ? AND password = ?", username, password).First(&data).Error
	//发送验证码
	err := utility.SendCode(data.Email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -201,
			"message": "验证码错误" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "验证码发送成功",
	})
}
