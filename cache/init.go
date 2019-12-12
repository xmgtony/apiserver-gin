package cache

import (
	"apidemo-gin/conf"
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)
var redisClient *redis.Client
// 初始化redisClient
func init()  {
	redisCfg := conf.Cfg.RedisCfg
	redisClient = redis.NewClient(&redis.Options{
		DB: redisCfg.Db,
		Addr: redisCfg.Addr,
		Password: redisCfg.Password,
		PoolSize: redisCfg.PoolSize,
		MinIdleConns: redisCfg.MinIdleConns,
		IdleTimeout: time.Duration(redisCfg.IdleTimeout) * time.Second,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("redis initialization failed: ",err)
	}
	log.Println("redis was initialized successfully")
}

func GetClient() *redis.Client {
	return redisClient
}