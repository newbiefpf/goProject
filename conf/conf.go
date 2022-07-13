package conf

/**
 * @Author newbie
 * @Date 22:11 2022/7/3
 **/
type Config struct {
	Mysql  mysqlConfig  `toml:"mysql"`
	Server serverConfig `toml:"server"`
}

type mysqlConfig struct {
	Username    string
	Password    string
	Dbname      string
	Host        string
	Port        int
	Protocol    string
	Charset     string
	MaxIdleConn int
	MaxOpenConn int
}

type serverConfig struct {
	Env  string
	Port int
}

var BlogConfig Config
