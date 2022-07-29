package models

import (
	//引入gorm
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goproject?charset=utf8mb4&parseTime=True&loc=Local"
	//gorm框架表名自动加s问题
	db, err := gorm.Open("mysql", dsn)
	//去掉自动添加的s
	db.SingularTable(true)

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Article{}, &ArticleInfo{}, &Attention{}, &CodeLite{})
	defer db.Close()
}
