package toml

import (
	"github.com/BurntSushi/toml"
	"time"
)

type Config struct {
	Title        string
	App          app
	File         file
	Jsonwebtoken jsonwebtoken
}

type app struct {
	Author  string
	Info    string
	Mark    string
	Release time.Time
}
type file struct {
	LogRouterPath  string `toml:"Log_ROUTER_PATH"`
	LogRouterName  string `toml:"LOG_ROUTER_NAME"`
	LogDetailsPath string `toml:"LOG_DETAILS_PATH"`
	LogDetailsName string `toml:"LOG_DETAILS_NAME"`
}
type jsonwebtoken struct {
	JwtSecret string
}

var ConfigToml Config

func ReadToml() {
	if _, err := toml.DecodeFile("./toml/example.toml", &ConfigToml); err != nil {
		panic(err)
	}

	//fmt.Printf("App信息：%+v\n\n", config.App)

}
