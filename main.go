package main

import (
	_ "apidemo-gin/cache"
	"apidemo-gin/conf"
	"apidemo-gin/router"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "config", "", "config file name")
	flag.Parse()
	// 加载配置文件
	if err := conf.LoadConfig(configFile); err != nil {
		panic(err)
	}
	// 设置gin启动模式，必须在创建gin实例之前
	gin.SetMode(conf.Cfg.Mode)
	// Create gin engine
	g := gin.New()
	// run mode
	var middlewares []gin.HandlerFunc

	router.Load(g, middlewares...)
	//Routes
	log.Printf("starting on port %s", conf.Cfg.Port)

	// health check
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("the server no response")
		}
		log.Println("the server started success!")
	}()
	// start server
	log.Printf(http.ListenAndServe(conf.Cfg.Port, g).Error())
}

// PingServer is be used to check server status
func pingServer() error {
	seconds := 1
	url := conf.Cfg.Url + conf.Cfg.Port + "/check/health"
	for i := 0; i < conf.Cfg.MaxPingCount; i++ {
		resp, err := http.Get(url)
		if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		log.Printf("waiting for the server online, sleep %d second", seconds)
		time.Sleep(time.Second * 1)
		seconds ++
	}
	return errors.New("Cannot connect to this server")
}
