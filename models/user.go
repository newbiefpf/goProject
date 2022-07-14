package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Userinfo struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(20);" json:"username"`
	Password string `gorm:"column:password;type:varchar(20);" json:"password"`
	Email    string `gorm:"column:email;type:varchar(36);" json:"email"`
	Age      int    `gorm:"column:age;type:int;" json:"age"`
	Sex      int    `gorm:"column:sex;type:int;" json:"sex"`
	Tel      string `gorm:"column:tel;type:varchar(20);" json:"tel"`
	Role     int    `gorm:"column:role;type:int;" json:"role"`
}

func (table *Userinfo) tableName() string {
	return "userinfo"
}
func GetUser() {
	userList := make([]*Userinfo, 0)
	DB.Find(&userList)

	for _, v := range userList {
		fmt.Printf("data=---->%v", v)
	}
}
