<p align="center">
        <img src="img/logo.jpg">
</p>

# bus
极简

让开发 golang 更轻松快速、不需要到处寻找资源库

    代码目录结构：
    assets          #
    config          #封装的 viper 包
    cors            #
    db              #数据库连接基础驱动包
    example         #数据库、makeapp、api、并发控制 etc. 的使用例子
    img             #
    logger          #封装的 zap 包
    msi             #Windows msi 编译包 ， 暂时未开放，需要使用请联系项目开发者shangguannihao@gmail.com
    pool            #
    xshell          #命令行窗口
    README.md       #说明书
目前分为如下几个部分：

    第一部分 标准化日志格式输出
            1)自定义配置文件路径(日常需要指定自己的配置文件路径
            2)日志统一（不需要使用API 让别人调用，只是普通的命令行程序打印日志
            3)api 日志统一（包括超时处理，跨域允许;使用API 的日志打印
    第二部分 数据库CRUD
            1)MySQL
              便捷DML、查询命令结果SHOW STATUS 、show slave status 、etc.
            2)Clickhouse
            3)努力更新中...
    第三部分 努力更新中...

# 第一部分 标准化日志格式输出
    #配置文件内容
    cat conf.yaml

    server:
      listen_ip: "0.0.0.0"
      listen_port: "80"
    log:
      path: "./agg.log"
      max_size: 100
      max_backup: 30
      max_age: 30
      level: "debug"

1)自定义配置文件路径(日常需要指定自己的配置文件路径，则引用格式如下)：

    # yaml 格式
    xxx.exe -conf=/xxx/yyy/zzz.yaml

    # json 格式
    xxx.exe -conf=/xxx/yyy/zzz.json

    # toml 格式
    xxx.exe -conf=/xxx/yyy/zzz.toml

    # ini 格式
    xxx.exe -conf=/xxx/yyy/zzz.ini

    # hcl 格式
    xxx.exe -conf=/xxx/yyy/zzz.hcl

2)日志统一（不需要使用API 让别人调用，只是普通的命令行程序，则引用格式如下：）

    #日志打印方式一（首先得有 yaml 配置文件）

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
api 日志统一（包括超时处理，跨域允许;使用API，则引用格式如下：）

    package main
    
    import (
        "fmt"
        "github.com/gin-contrib/cors"
        "github.com/gin-gonic/gin"
        "github.com/org-lib/bus/config"
        "github.com/org-lib/bus/handler"
        "github.com/org-lib/bus/logger"
        "github.com/vearne/gin-timeout"
        "go.uber.org/zap"
        "net/http"
        "runtime"
        "time"
    )
    var (
        router = gin.Default()
        defaultMsg = `{"code": -1, "msg":"http: Handler timeout"}`
        MaxProces = runtime.NumCPU()
    )
    
    //允许跨域
    func Cors() gin.HandlerFunc {
        handlerFunc := cors.New(cors.Config{
            AllowMethods:     []string{"*"},
            AllowHeaders:     []string{"Authentication"}, //此处设置非默认之外的请求头(自定义请求头),否则会出现跨域问题
            AllowAllOrigins:  true,
            AllowCredentials: true,
            MaxAge:           24 * time.Hour,
        })
        return handlerFunc
    }
    
    func main() {
        //并发能力控制
    
        if MaxProces > 2 {
            MaxProces -= 1
        }
        runtime.GOMAXPROCS(MaxProces)
    
        // 设置gin启动模式为生产模式
    
        gin.SetMode(gin.ReleaseMode)
    
        //跨域
        router.Use(Cors())
    
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
            v1.GET("/hello", handler.HelloWorld)
    
        }
        fmt.Println()
    
        // 启动服务，获取配置文件config.yaml的IP和端口：listen_ip和listen_port
    
        addr := fmt.Sprintf("%v:%v",config.Config.V.GetString("server.listen_ip") ,config.Config.V.GetString("server.listen_port"))
    
        //输出一个名为message的自定义内容值、{"message":"Start server"}，以及自定义key：value 的输出
        //{"level":"INFO","timestamp":"2021-12-22 13:38:09:000","caller":"example/main.go:68","message":"Start server","listen":"0.0.0.0:80"}
    
        logger.Log.Info("Start server", zap.String("listen", addr))
        err := router.Run(fmt.Sprintf("%v",addr))
        if err != nil {
            logger.Log.Error("Start server", zap.String("error", err.Error()))
        }
        //logger.Log.Info("Start server success", zap.String("listen", addr))
    
    }
# 第二部分 数据库CRUD
1)MySQL
请参考 example/mysql

2)Clickhouse
请参考 example/clickhouse
