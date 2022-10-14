package main

import (
	"fmt"
	"github.com/org-lib/bus/consul"
)

func main() {
	info := &consul.ClientInfo{
		Name:    "serverNode",
		Tag:     "v1000",
		Address: "localhost:8500",
	}
	mp, err := consul.SearchServer(info)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(mp)
}
