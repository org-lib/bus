package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/org-lib/bus/config"
	"github.com/org-lib/bus/cors"
	"github.com/org-lib/bus/example/concurrency/handler"
	"github.com/org-lib/bus/logger"
	"github.com/org-lib/bus/pool"
	"github.com/vearne/gin-timeout"
	"go.uber.org/zap"
	"net/http"
	"runtime"
	"time"
)

var (
	router     = gin.Default()
	defaultMsg = `{"code": -1, "msg":"http: Handler timeout"}`
	MaxProces  = runtime.NumCPU()
)

func main1() {
	//并发能力控制

	if MaxProces > 2 {
		MaxProces -= 1
	}
	runtime.GOMAXPROCS(MaxProces)

	// 设置gin启动模式为生产模式

	gin.SetMode(gin.ReleaseMode)

	//跨域
	router.Use(cors.Cors())

	router.Use(timeout.Timeout(
		timeout.WithTimeout(20*time.Second),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout), // optional
		timeout.WithDefaultMsg(defaultMsg),                   // optional
		timeout.WithCallBack(func(r *http.Request) {
			fmt.Println("timeout happen, url:", r.URL.String())
		}))) // optional

	router.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 管理API
	v1 := router.Group("api")
	{
		v1.GET("/add", handler.Add)
		v1.GET("/del", handler.Del)
	}

	// 启动服务，获取配置文件config.yaml的IP和端口：listen_ip和listen_port

	addr := fmt.Sprintf("%v:%v", config.Config.V.GetString("server.listen_ip"), config.Config.V.GetString("server.listen_port"))

	//输出一个名为message的自定义内容值、{"message":"Start server"}，以及自定义key：value 的输出
	//{"level":"INFO","timestamp":"2021-12-22 13:38:09:000","caller":"example/main.go:68","message":"Start server","listen":"0.0.0.0:80"}

	logger.Log.Info("Start server", zap.String("listen", addr))
	err := router.Run(fmt.Sprintf("%v", addr))
	if err != nil {
		logger.Log.Error("Start server", zap.String("error", err.Error()))
	}
	//logger.Log.Info("Start server success", zap.String("listen", addr))

}
func main2() {
	//并发能力控制

	if MaxProces > 2 {
		MaxProces -= 1
	}
	runtime.GOMAXPROCS(MaxProces)

	// 设置gin启动模式为生产模式

	gin.SetMode(gin.ReleaseMode)

	//跨域
	router.Use(cors.Cors())

	router.Use(timeout.Timeout(
		timeout.WithTimeout(20*time.Second),
		timeout.WithErrorHttpCode(http.StatusRequestTimeout), // optional
		timeout.WithDefaultMsg(defaultMsg),                   // optional
		timeout.WithCallBack(func(r *http.Request) {
			fmt.Println("timeout happen, url:", r.URL.String())
		}))) // optional

	router.Use(logger.GinLogger(), logger.GinRecovery(true))

	config.Work = pool.NewPool(config.Config.V.GetInt("pool.max"))
	// 管理API
	v1 := router.Group("api")
	{
		v1.GET("/add", handler.AddPool)
		v1.GET("/del", handler.DelPool)

	}
	config.Work.Wait()

	// 启动服务，获取配置文件config.yaml的IP和端口：listen_ip和listen_port

	addr := fmt.Sprintf("%v:%v", config.Config.V.GetString("server.listen_ip"), config.Config.V.GetString("server.listen_port"))

	//输出一个名为message的自定义内容值、{"message":"Start server"}，以及自定义key：value 的输出
	//{"level":"INFO","timestamp":"2021-12-22 13:38:09:000","caller":"example/main.go:68","message":"Start server","listen":"0.0.0.0:80"}

	logger.Log.Info("Start server", zap.String("listen", addr))
	err := router.Run(fmt.Sprintf("%v", addr))
	if err != nil {
		logger.Log.Error("Start server", zap.String("error", err.Error()))
	}
	//logger.Log.Info("Start server success", zap.String("listen", addr))

}
func main() {
	//main1()
	main2()
}
