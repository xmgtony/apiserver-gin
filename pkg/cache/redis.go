package cache

import (
	"apiserver-gin/pkg/config"
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)

var redisClient *redis.Client

// 初始化redisClient
func RedisInit() {
	redisCfg := config.Cfg.RedisCfg
	redisClient = redis.NewClient(&redis.Options{
		DB:           redisCfg.Db,
		Addr:         redisCfg.Addr,
		Password:     redisCfg.Password,
		PoolSize:     redisCfg.PoolSize,
		MinIdleConns: redisCfg.MinIdleConns,
		IdleTimeout:  time.Duration(redisCfg.IdleTimeout) * time.Second,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("redis initialization failed: ", err)
	}
	log.Println("redis was initialized successfully")
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func RedisClose() {
	_ = redisClient.Close()
}
