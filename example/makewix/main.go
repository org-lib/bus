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

	//svc  			定义服务显示名称，
	//name 			定义产品名称以及exe程序包名称，
	//filetype 		定义 json 内容
	//二进制可执行文件 需要放在当前目录build/amd64/bus.exe
	//启动的配置文件	需要放在当前目录assets/config.yaml
	//ico.ico		需要放在当前目录assets/ico.ico

	msi.SetJson(os.Args[1], os.Args[2], os.Args[3], wix)
}
