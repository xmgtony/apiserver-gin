package model

import (
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/db"
)

var dataBase *db.DataBase

func InitDB(c config.DataBaseConfig) {
	dataBase = db.New(c)
}

func CloseDb() {
	dataBase.Close()
}
