package aboutUser

import (
	"fmt"
	"goProject/db"
	"goProject/model/dbModel"
	"goProject/model/tableModel"
	"goProject/utility"
)

func AddUser(u dbModel.User) {
	var DB = db.ConnectDb
	DB.Create(&u)
}
func FindUser(u tableModel.User) (token string) {
	var DB = db.ConnectDb
	var duser dbModel.User

	result := DB.Where("username=?  AND username=?", u.Username, u.Password).First(&duser)
	fmt.Println(duser)
	fmt.Println(result.Error)
	if result.Error != nil {
		panic("账号密码错误")
	}
	token, err := utility.GenerateToken(u.Username, u.Password)
	if err != nil {

		panic("令牌生成错误")
	}
	return token
}
