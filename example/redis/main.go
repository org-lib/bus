package main

import (
	"fmt"
	"github.com/org-lib/bus/config"
	"github.com/org-lib/bus/db/redis"
	"github.com/org-lib/bus/logger"
)

func main() {
	//定义 cfg 对象
	var cfg *redis.Info
	cfg = &redis.Info{
		Host:     config.Config.V.GetString("redis.host"),
		Port:     config.Config.V.GetInt("redis.port"),
		DB:       config.Config.V.GetInt("redis.db"),
		Password: config.Config.V.GetString("redis.password"),
	}
	client := redis.NewClient(cfg)
	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Redis connection fail:%v, pong:%v", err, pong))
	}
	fmt.Println(client.Info().Result())
	logger.Log.Info(fmt.Sprintf("Redis connection success!"))

}
