package redis

import (
	"fmt"
	"github.com/go-redis/redis"
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
