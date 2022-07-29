package main

import (
	dbmodel "goProject/models"
	"goProject/router"
)

// @title 接口文档
// @version 1.0
// @description 笔记的后台api
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8888
// @BasePath /
func main() {
	dbmodel.Init()
	router.Router()

}
