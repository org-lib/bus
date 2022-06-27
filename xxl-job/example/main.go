package main

import (
	"github.com/org-lib/bus/xxl-job"
	"github.com/xxl-job/xxl-job-executor-go/example/task"
	"log"
	"time"
)

func main() {
	// 初始化执行器
	job := &xxl.Jobs{
		XxlServer:   "http://devops-xxl-job-admin-sit.cloud.bz/xxl-job-admin",
		Token:       "",
		Port:        "9998",
		RegistryKey: "ydq-jobs",
		MaxAge:      2 * time.Second,
	}
	// 初始化 gin 路由
	job.Init()
	defer job.Stop()

	// 注册任务 handler
	job.RegTask("task.test", task.Test)
	job.RegTask("task.test2", task.Test2)
	job.RegTask("task.panic", task.Panic)

	log.Fatal(job.Run())
}
