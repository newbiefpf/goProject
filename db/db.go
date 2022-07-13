package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goProject/conf"
	"log"
)

var DB *gorm.DB

func InitDB() {
	//配置连接数据库
	mysqlConf := conf.BlogConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@/%s?charset=%s", mysqlConf.Username, mysqlConf.Password, mysqlConf.Dbname, mysqlConf.Charset)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("open db error: %s", err.Error())
	}
	db.DB().SetMaxIdleConns(mysqlConf.MaxIdleConn)
	db.DB().SetMaxOpenConns(mysqlConf.MaxOpenConn)
	db.SingularTable(true)
	db.LogMode(true)

	if err = db.DB().Ping(); err != nil {
		log.Printf("ping error: %s", err.Error())
	}
	DB = db
}
