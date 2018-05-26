package config

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	cfg *ini.File
)

type dbConfig struct {
	Dialect  string
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

type appConfig struct {
	PageSize    string
	TokenSecret string
}
type serverConfig struct {
	Port      int
	Env       string
	APIPrefix string
}

func init() {
	var err error
	cfg, err = ini.Load("./app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	loadDBConifg()
	loadAppConfig()
	loadServerConfig()
}

// DBConfig 数据库相关配置
var DBConfig dbConfig

func loadDBConifg() {
	err := cfg.Section("database").MapTo(&DBConfig)
	if err != nil {
		log.Fatalf("Fail to load database config: %v", err)
	}
}

//AppConfig 服务内部相关配置
var AppConfig appConfig

func loadAppConfig() {
	err := cfg.Section("app").MapTo(&DBConfig)
	if err != nil {
		log.Fatalf("Fail to load app config: %v", err)
	}
}

//ServerConfig 服务运行环境配置
var ServerConfig serverConfig

func loadServerConfig() {
	err := cfg.Section("server").MapTo(&DBConfig)
	if err != nil {
		log.Fatalf("Fail to load server config: %v", err)
	}
}
