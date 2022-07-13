package main

/**
 * @Author newbie
 * @Date 22:11 2022/7/3
 **/
import (
	"fmt"
	"goProject/conf"
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
	log.Println("Hello newbie!") // log 还是可以作为输出的前缀
	if err := conf.Init(); err != nil {
		log.Printf("conf.Init() err:%+v", err)
	}
	mysqlConf := conf.Conf.Mysql
	fmt.Println(mysqlConf.DbName)
}
