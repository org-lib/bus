package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"strings"
)

type Info struct {
	Host     string
	Port     int
	DB       int
	Password string
}

func NewClient(cnf *Info) *redis.Client {
	if cnf.Host == "" {
		fmt.Println("使用了默认地址：localhost")
		cnf.Host = "localhost"
	}
	if cnf.Port < 0 {
		fmt.Println("使用了默认端口：6379")
		cnf.Port = 6379
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%d", cnf.Host, cnf.Port),
		Password: cnf.Password,
		DB:       cnf.DB,
	})
	return client
}

func InfoToMap(client *redis.Client) (map[string]map[string]interface{}, error) {
	msg, err := client.Info().Result()
	if err != nil {
		return nil, err
	}
	info := make(map[string]map[string]interface{})
	subInfo := make(map[string]interface{})

	// one end

	info_title := ""
	for _, s := range strings.Split(string(msg), "\r\n") {
		if strings.Trim(s, " ") == "" {
			continue
		}
		if strings.HasPrefix(s, "# ") {
			if info_title != "" {
				info[info_title] = subInfo

				//重置子 map

				subInfo = nil
				subInfo = make(map[string]interface{})
			}
			info_title = strings.ReplaceAll(s, "# ", "")
			continue
		}
		kv := strings.Split(s, ":")
		subInfo[kv[0]] = kv[1]
	}
	// 处理最后一个map，除非Redis info 命令结果集，最后一行不是空行或者nil的标识

	info[info_title] = subInfo

	//重置子 map

	subInfo = nil

	info_title = ""

	return info, nil
}
