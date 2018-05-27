package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/zzayne/go-blog/config"
)

var DB *gorm.DB

func initDB() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.Url)
	if err != nil {
		fmt.Println(err.Error())
	}

	if config.ServerConfig.Env == DevelopmentMode {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	DB = db

}

func init() {
	initDB()
}
