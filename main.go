package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goProject/api"
	dbContent "goProject/db"
	_ "goProject/docs"
	"goProject/toml"
)

func main() {
	toml.ReadToml()
	dbContent.Init()
	api.Init()
	//数据库
}
