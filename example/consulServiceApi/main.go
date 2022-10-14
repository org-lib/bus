package main

import (
	"fmt"
	"github.com/org-lib/bus/consul"
)

func main() {
	// 注册一个名字为 serverNode 的服务到 consul server，以便于调用 我的ip和端口
	info := &consul.Info{
		ID:                             "serverNode_1",
		Name:                           "serverNode",
		Port:                           9527,
		Tags:                           []string{"v1000"},
		Address:                        consul.LocalIP(),
		ConsulAddress:                  "localhost:8500",
		CheckPort:                      8080,
		CheckTimeout:                   "3s",
		CheckInterval:                  "5s",
		DeregisterCriticalServiceAfter: "30s",
	}
	err := consul.RegisterServer(info)
	if err != nil {
		fmt.Println(err.Error())
	}
}
