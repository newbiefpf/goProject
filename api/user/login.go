package user

import (
	"github.com/gin-gonic/gin"
	"goProject/db"
	"goProject/model/dbModel"
	"goProject/model/tableModel"
	"goProject/utility"
	"goProject/utility/ginResult"
)

// FindUserByName 查询用户信息
// @Summary 查询用户信息
// @Description 根据用户名查询用户信息
// @Tags 用户接口
// @Accept application/json
// @Produce application/json
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /Login [post]
func Login(c *gin.Context) {
	var user tableModel.User
	var duser dbModel.User
	_ = c.BindJSON(&user)
	if user.Username == "" {
		c.JSON(201, ginResult.ErrParam.WithMsg("账号不能为空"))
		return
	}
	if user.Password == "" {
		c.JSON(201, ginResult.ErrParam.WithMsg("密码不能为控"))
		return
	}
	var DB = db.ConnectDb
	result := DB.Where("username=?  AND password=?", user.Username, user.Password).First(&duser)
	if result.Error != nil {
		c.JSON(200, ginResult.ErrParam.WithMsg("账号密码错误"))
		return
	}
	token, err := utility.GenerateToken(user.Username, user.Password)
	if err != nil {
		c.JSON(200, ginResult.ErrSignParam)
		return
	}
	res := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}
	c.JSON(200, ginResult.OK.WithData(res))
}
