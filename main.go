package main

import (
	"apiserver-gin/internal/handler"
	"apiserver-gin/internal/model"
	"apiserver-gin/internal/router"
	"apiserver-gin/internal/service"
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/log"
	"apiserver-gin/pkg/version"
	"apiserver-gin/server"
)

func init() {
	// 解析服务器启动参数
	appOpt := &server.AppOptions{}
	server.ResolveAppOptions(appOpt)
	if appOpt.PrintVersion {
		version.PrintVersion()
	}
	// 加载配置文件
	c := config.Load(appOpt.ConfigFilePath)
	log.LoggerInit(&c.LogConfig, c.AppName) // 日志
	model.InitDB(c.DbConfig)                // 数据库连接初始化
	s := service.InitService()              // service初始化
	handler.InitHandler(s)                  // handler层注入service
}

func main() {
	srv := server.NewHttpServer(config.GlobalConfig)
	srv.RegisterOnShutdown(func() {
		model.CloseDb()
	})
	srv.Run(router.Load)
}
