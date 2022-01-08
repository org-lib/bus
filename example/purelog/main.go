package main

import (
	"github.com/org-lib/bus/logger"
	"go.uber.org/zap"
)

func main() {

	//输出一个名为message的自定义内容值、{"message":"Start server"}，以及自定义 key：value 的输出
	//{"level":"INFO","timestamp":"2021-12-22 13:38:09:000","caller":"example/main.go:14","message":"Start server","listen":"0.0.0.0:33333"}

	logger.Log.Info("server start", zap.String("listen", "33333"))
}
