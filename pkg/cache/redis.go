package cache

import (
	"apiserver-gin/pkg/config"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisClient *redis.Client

// 初始化redisClient
func RedisInit(config config.Config) {
	redisCfg := config.RedisCfg
	redisClient = redis.NewClient(&redis.Options{
		DB:           redisCfg.Db,
		Addr:         redisCfg.Addr,
		Password:     redisCfg.Password,
		PoolSize:     redisCfg.PoolSize,
		MinIdleConns: redisCfg.MinIdleConns,
		IdleTimeout:  time.Duration(redisCfg.IdleTimeout) * time.Second,
	})
	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func RedisClose() {
	_ = redisClient.Close()
}
