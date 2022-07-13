package conf

/**
 * @Author newbie
 * @Date 22:11 2022/7/3
 **/
import (
	"flag"
	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf global
	Conf = &Config{}
)

// Config .
type Config struct {
	Mysql *Mysql
}

type Mysql struct {
	UserName string
	Password string
	IpHost   string
	DbName   string
}

func init() {
	flag.StringVar(&confPath, "conf", "./conf/conf.toml", "-conf path")
}

// Init init conf
func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
