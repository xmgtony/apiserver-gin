// Created on 2021/5/4.
// @author tony
// email xmgtony@gmail.com
// description 配置mysql链接

package db

import (
	"apiserver-gin/pkg/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Dao 数据访问对象
type Dao struct {
	*gorm.DB
}

func New(c config.DataBaseConfig) *Dao {
	return &Dao{
		openDB(
			c.Username,
			c.Password,
			c.Host,
			c.Port,
			c.Dbname,
			c.MaximumPoolSize,
			c.MaximumIdleSize),
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
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetConnMaxLifetime(time.Hour)
	// 设置连接池大小
	sqlDb.SetMaxOpenConns(maxPoolSize)
	sqlDb.SetMaxIdleConns(maxIdle)
	return db
}

func (d *Dao) Close() {
	db, _ := d.DB.DB()
	_ = db.Close()
}
