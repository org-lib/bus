package main

import (
	"github.com/org-lib/bus/msi"
	"os"
)

var (
	wix = "wix.json"
)

func main() {
	/*
		创建一个名为 wix.json 的文件
	*/

	//os.Args[1] svc  			定义服务显示名称，
	//os.Args[2] name 			定义产品名称以及exe程序包名称，
	//os.Args[3] filetype 		文件 json 类型
	//os.Args[4] conf	 		定义 配置文件路径名称
	//os.Args[5] wix.json	 		定义 wix.json
	//二进制可执行文件 需要放在当前目录build/amd64/bus.exe
	//启动的配置文件	需要放在当前目录assets/config.yaml
	//ico.ico		需要放在当前目录assets/ico.ico

	msi.SetJson(os.Args[1], os.Args[2], os.Args[3], os.Args[4], wix)
}
