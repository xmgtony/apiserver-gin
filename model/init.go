package model

import (
	"apidemo-gin/pkg/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DataBase 用来组织数据库信息，实际使用中可能会有Master和Slave主从库
type DataBase struct {
	Master *gorm.DB
}

var DB *DataBase

func DBInit() {
	cfg := config.Cfg.Database
	DB = &DataBase{
		Master: openDB(cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname),
	}
}

func openDB(user, password, host, port, dbname string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// 设置日志格式和连接池大小
	db.LogMode(config.Cfg.Database.LogMode)
	db.DB().SetMaxOpenConns(config.Cfg.Database.MaximumPoolSize)
	db.DB().SetMaxIdleConns(config.Cfg.Database.MaximumIdleSize)

	return db
}

func DBClose() {
	_ = DB.Master.Close()
}
