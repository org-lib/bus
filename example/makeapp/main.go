package main

import (
	"fmt"
	"github.com/org-lib/bus/example/makeapp/app"
	"github.com/org-lib/bus/logger"
	"github.com/org-lib/bus/msi"
	"go.uber.org/zap"
)

func main() {
	/*
		创建一个名为 Bus.msi 的安装包
	*/

	//全局异常抓捕
	E()

	//定义 msi 显示的服务名称 "Bus"
	//msi.Bus 可以是任意可以在后台运行的api 服务，或者其他任意后台运行的服务
	//如果不是能后台长期运行的服务，一次性运行结束的程序有什么意义？没意义
	run := app.Bus
	msi.Svc("Bus", run)
}
func E() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("*** 异常:", err)
			logger.Log.Error("*** 异常:", zap.String("remote", fmt.Sprintf("%v", err)))
		}
	}()
}
