package db

import (
	"fmt"
	//引入gorm
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ConnectDb *gorm.DB

func Init() {
	//连接数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/goproject?charset=utf8mb4&parseTime=True&loc=Local"
	//gorm框架表名自动加s问题
	db, err := gorm.Open("mysql", dsn)
	//去掉自动添加的s
	db.SingularTable(true)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	//创建数据库中的标
	//var user dbModel.User
	//var article dbModel.Article
	//db.AutoMigrate(&user, &article)
	ConnectDb = db
	//defer db.Close()
}
