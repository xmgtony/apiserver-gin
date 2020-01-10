package main

import (
	"apidemo-gin/cache"
	"apidemo-gin/model"
	"apidemo-gin/pkg/config"
	"apidemo-gin/pkg/log"
	"apidemo-gin/pkg/version"
	"apidemo-gin/router"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

// Profile name, no suffix required, default "config"
// For example, "config" stands for config.yml
var configFile string

func main() {
	var printVersion bool
	flag.StringVar(&configFile, "config", "", "config file name")
	flag.BoolVar(&printVersion,
		"v",
		false,
		"-v be used to print app version info",
	)
	flag.Parse()
	if printVersion {
		version.PrintVersion()
	}
	// 加载配置文件
	config.Load(configFile)
	// 初始化Redis Client
	cache.RedisInit()
	defer cache.RedisClose()
	// 初始化数据库信息
	model.DBInit()
	defer model.DBClose()
	// 初始化logger
	log.LoggerInit()
	// 便于在外部挂载middleware，添加到当前slice中即可
	var middlewares []gin.HandlerFunc
	// 设置gin启动模式，必须在创建gin实例之前
	gin.SetMode(config.Cfg.Mode)
	// Create gin engine
	g := gin.New()
	router.Load(g, middlewares...)
	// health check
	go func() {
		if err := ping(); err != nil {
			log.Log.Fatal("the server no response")
		}
		log.Log.Info("the server started success!")
	}()
	// Start on the specified port
	log.Log.Info(g.Run(config.Cfg.Port).Error())
}

// PingServer is be used to check server status
func ping() error {
	seconds := 1
	url := config.Cfg.Url + config.Cfg.Port + "/check/health"
	for i := 0; i < config.Cfg.MaxPingCount; i++ {
		resp, err := http.Get(url)
		if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		log.Log.Info(fmt.Sprintf("waiting for the server online, sleep %d second", seconds))
		time.Sleep(time.Second * 1)
		seconds++
	}
	return errors.New(fmt.Sprintf("Can not connect to this server on port %s", config.Cfg.Port))
}
