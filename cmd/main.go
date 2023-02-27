package main

import (
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/db"
	"apiserver-gin/pkg/log"
	"apiserver-gin/pkg/version"
	"apiserver-gin/server"
)

func main() {
	// 解析服务器启动参数
	appOpt := &server.AppOptions{}
	server.ResolveAppOptions(appOpt)
	if appOpt.PrintVersion {
		version.PrintVersion()
	}
	// 加载配置文件
	c := config.Load(appOpt.ConfigFilePath)
	log.InitLogger(&c.LogConfig, c.AppName) // 日志
	ds := db.NewDefaultMysql(c.DBConfig)    // 创建数据库链接，使用默认的实现方式
	// 创建HTTPServer
	srv := server.NewHttpServer(config.GlobalConfig)
	srv.RegisterOnShutdown(func() {
		if ds != nil {
			ds.Close()
		}
	})
	routers := getRouters(ds)
	srv.Run(routers...)
}
