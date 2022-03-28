package main

import (
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/db"
	"apiserver-gin/pkg/log"
	"apiserver-gin/pkg/validator"
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
	validator.LazyInitGinValidator(c.Language) // gin validator自定义
	log.InitLogger(&c.LogConfig, c.AppName)    // 日志
	ds := db.NewDefaultMysql(c.DBConfig)       // 创建数据库链接，使用默认的实现方式
	// 依赖较多时可以拆分出去，使用wire解决依赖关系
	//userRepo := mysql.NewUserRepo(ds)
	//userSrv := service.NewUserService(userRepo) // 创建userService
	//userHandler := user.NewUserHandler(userSrv) // 创建userHandler
	//rt := router.NewRouter(userHandler)         // router 包注入userHandler
	routers := getRouters(ds)
	// 创建HTTPServer
	srv := server.NewHttpServer(config.GlobalConfig)
	srv.RegisterOnShutdown(func() {
		if ds != nil {
			ds.Close()
		}
	})
	srv.Run(routers...)
}
