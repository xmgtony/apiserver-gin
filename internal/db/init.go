package db

import (
	"apiserver-gin/pkg/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// DataBase 用来组织数据库信息，实际使用中可能会有Master和Slave主从库
type DataBase struct {
	Master *gorm.DB
}

var DB *DataBase

func DataBaseInit(config config.Config) {
	dbConfig := config.Database
	DB = &DataBase{
		Master: openDB(
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Dbname,
			dbConfig.MaximumPoolSize,
			dbConfig.MaximumIdleSize),
	}
}

func openDB(user, password, host, port, dbname string, maxPoolSize, maxIdle int) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // 缓存每一条sql语句，提高执行速度
	})
	if err != nil {
		panic(err)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetConnMaxLifetime(time.Hour)
	// 设置连接池大小
	sqlDb.SetMaxOpenConns(maxPoolSize)
	sqlDb.SetMaxIdleConns(maxIdle)
	return db
}

func DBClose() {
	sqlDb, _ := DB.Master.DB()
	_ = sqlDb.Close()
}
