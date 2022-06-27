package xxl

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-middleware/xxl-job-executor"
	"github.com/xxl-job/xxl-job-executor-go"
	"sync"
	"time"
)

var (
	duplicate map[string]string
	lock      sync.Mutex
	xxlRouter *gin.Engine
	exec      xxl.Executor
)

type Jobs struct {
	XxlServer   string        `json:"xxlServer"` // XxlServer 指定一个xxl-job 的调度地址
	Token       string        `json:"token"`
	Port        string        `json:"port"` // Port RegistryKey ，同一个服务，则同一个端口和注册名 RegistryKey；
	RegistryKey string        `json:"registryKey"`
	MaxAge      time.Duration `json:"maxAge"` // gin router 请求最大时长
}

func (job *Jobs) Init() {
	// 跨域配置
	handlerFunc := cors.New(cors.Config{
		//AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Authentication"}, //此处设置非默认之外的请求头(自定义请求头),否则会出现跨域问题
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           job.MaxAge,
	})

	//初始化执行器
	exec = xxl.NewExecutor(
		xxl.ServerAddr(job.XxlServer),
		xxl.AccessToken(job.Token), //请求令牌(默认为空)
		//xxl.ExecutorIp("127.0.0.1"),      //可自动获取
		xxl.ExecutorPort(job.Port),       //默认9999（此处要与gin服务启动port必需一至）
		xxl.RegistryKey(job.RegistryKey), //执行器名称
	)
	exec.Init()

	// 初始化gin路由
	xxlRouter = gin.Default()
	xxlRouter.Use(handlerFunc)
	xxl_job_executor_gin.XxlJobMux(xxlRouter, exec)

	//注册gin的handler
	xxlRouter.GET("ping", func(cxt *gin.Context) {
		cxt.JSON(200, "pong")
	})
}

func (job *Jobs) Run() (err error) {
	return xxlRouter.Run(":" + job.Port)
}

func (job *Jobs) RegTask(pattern string, task xxl.TaskFunc) {
	//注册任务handler
	exec.RegTask(pattern, task)
}

func (job *Jobs) Stop() {
	exec.Stop()
}
