package main

import (
	"apiserver-gin/internal/handler/v1/user"
	"apiserver-gin/internal/repo/mysql"
	"apiserver-gin/internal/router"
	"apiserver-gin/internal/service"
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
	// 依赖较多时可以拆分出去，使用wire解决依赖关系
	ds := db.NewDefaultMysql(c.DBConfig) // 创建数据库链接，使用默认的实现方式
	userRepo := mysql.NewUserRepo(ds)
	userSrv := service.NewUserService(userRepo) // 创建userService
	userHandler := user.NewUserHandler(userSrv) // 创建userHandler
	rt := router.NewRouter(userHandler)         // router 包注入userHandler

	// 创建HTTPServer
	srv := server.NewHttpServer(config.GlobalConfig)
	srv.RegisterOnShutdown(func() {
		if ds != nil {
			ds.Close()
		}
	})
	srv.Run(rt)
}
