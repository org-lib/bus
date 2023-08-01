package main

import (
	"fmt"
	"github.com/org-lib/bus/config"
	"github.com/org-lib/bus/db/redis"
	"github.com/org-lib/bus/logger"
)

func main() {
	//定义 cfg 对象
	var cfg redis.Info
	cfg = redis.Info{
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

	var redisInfo = make(map[string]map[string]interface{})
	redisInfo, err = redis.InfoToMap(client)
	//_ = redisInfo
	//fmt.Println(client.Info().String())
	fmt.Println("*********************************** CPU")
	fmt.Println(redisInfo["CPU"])
	fmt.Println(client.Info("CPU").String())
	fmt.Println("*********************************** Server")
	fmt.Println(redisInfo["Server"])
	fmt.Println(client.Info("Server").String())
	fmt.Println("*********************************** Clients")
	fmt.Println(redisInfo["Clients"])
	fmt.Println(client.Info("Clients").String())
	fmt.Println("*********************************** Memory")
	fmt.Println(redisInfo["Memory"])
	fmt.Println(client.Info("Memory").String())
	fmt.Println("*********************************** Persistence")
	fmt.Println(redisInfo["Persistence"])
	fmt.Println(client.Info("Persistence").String())
	fmt.Println("*********************************** Stats")
	fmt.Println(redisInfo["Stats"])
	fmt.Println(client.Info("Stats").String())
	fmt.Println("*********************************** Replication")
	fmt.Println(redisInfo["Replication"])
	fmt.Println(client.Info("Replication").String())
	fmt.Println("*********************************** Cluster")
	fmt.Println(redisInfo["Cluster"])
	fmt.Println(client.Info("Cluster").String())
	fmt.Println("*********************************** Keyspace")
	fmt.Println(redisInfo["Keyspace"])
	fmt.Println(client.Info("Keyspace").String())

	//info2map
	var redisInfo2 = make(map[string]interface{})
	redisInfo2, err = redis.Info2Map(client)

	fmt.Println(redisInfo2["role"])
	logger.Log.Info(fmt.Sprintf("Redis connection success!"))

}
