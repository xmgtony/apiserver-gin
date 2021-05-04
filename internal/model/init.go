package model

import (
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/db"
)

var dao *db.Dao

func InitDB(c config.DataBaseConfig) {
	dao = db.New(c)
}

func CloseDb() {
	dao.Close()
}
