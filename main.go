package main

/**
 * @Author newbie
 * @Date 22:11 2022/7/3
 **/
import (
	"flag"
	"github.com/BurntSushi/toml"
	"goProject/apis"
	"goProject/conf"
	"goProject/db"
	"log"
	"os"
)

func init() {
	file := "./logInfo/" + "message" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Printf("初始化日志文件 err:%+v", err)
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[logFileTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
func main() {
	log.Println("Hello newbie!") // log
	//读取配置文件
	var configFile string
	flag.StringVar(&configFile, "c", "./conf/conf.toml", "config file")
	log.Printf("读取配置文件: %s ", configFile)
	if _, err := toml.DecodeFile(configFile, &conf.BlogConfig); err != nil {
		log.Printf("Failed to load config file: %s err: %s ", configFile, err.Error())
	}
	db.InitDB()
	defer db.DB.Close()
	//项目启动
	log.Println("项目启动")
	apis.Route()
}
