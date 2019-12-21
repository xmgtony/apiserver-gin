package main

import (
	"apidemo-gin/cache"
	"apidemo-gin/conf"
	"apidemo-gin/model"
	"apidemo-gin/router"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

// Profile name, no suffix required, default "config"
// For example, "config" stands for config.yml
var configFile string

func main() {
	flag.StringVar(&configFile, "config", "", "config file name")
	flag.Parse()
	// 加载配置文件
	conf.Load(configFile)
	// 初始化Redis Client
	cache.RedisInit()
	defer cache.RedisClose()
	// 初始化数据库信息
	model.DBInit()
	defer model.DBClose()

	// 便于在外部挂载middleware，添加到当前slice中即可
	var middlewares []gin.HandlerFunc
	// 设置gin启动模式，必须在创建gin实例之前
	gin.SetMode(conf.Cfg.Mode)
	// Create gin engine
	g := gin.New()
	router.Load(g, middlewares...)
	//Routes
	log.Printf("start up success on port %s", conf.Cfg.Port)

	// health check
	go func() {
		if err := ping(); err != nil {
			log.Fatal("the server no response")
		}
		log.Println("the server started success!")
	}()
	// Start on the specified port
	log.Printf(g.Run(conf.Cfg.Port).Error())
}

// PingServer is be used to check server status
func ping() error {
	seconds := 1
	url := conf.Cfg.Url + conf.Cfg.Port + "/check/health"
	for i := 0; i < conf.Cfg.MaxPingCount; i++ {
		resp, err := http.Get(url)
		if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		log.Printf("waiting for the server online, sleep %d second", seconds)
		time.Sleep(time.Second * 1)
		seconds++
	}
	return errors.New(fmt.Sprintf("Can not connect to this server on port %s", conf.Cfg.Port))
}
