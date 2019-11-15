package app

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName     string `envconfig:"APP_NAME" default:"golang-gin-sample"`
	AppVersion  string `envconfig:"APP_VERSION" default:"undefined"`
	RdbType     string `envconfig:"RDB_TYPE" default:"mysql"`
	RdbUser     string `envconfig:"RDB_USER" default:"testuser"`
	RdbPassword string `envconfig:"RDB_PASSWORD" default:"testpass"`
	RdbProtocol string `envconfig:"RDB_PROTOCOL" default:"tcp"`
	RdbHost     string `envconfig:"RDB_HOST" default:"127.0.0.1:3306"`
	RdbName     string `envconfig:"RDB_NAME" default:"testdb"`
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	InitConfig()
	return config
}

func InitConfig() {
	once.Do(func() {
		config = &Config{}
		envconfig.Process("", config)
	})
}
