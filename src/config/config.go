package config

import "github.com/BurntSushi/toml"

var (
	C Cfg
)

type Cfg struct {
	Port      string
	Endpoints []string
}

func Init() {
	if _, err := toml.DecodeFile("./conf/config.toml", &C); err != nil {
		panic(err)
	}
}
