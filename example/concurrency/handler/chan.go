package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	ch = make(chan bool, 1)
)

func Add(ctx *gin.Context) {
	//不使用 goto 标签，是因为我是同步的请求响应，如果是要求异步，且会不断推送状态到前端，则可以用 goto
	//顾，返回超时退出后，前端重新尝试即可
	select {
	case ch <- true:
		fmt.Println("开始执行任务..")
		break
	case <-time.After(10 * time.Second):
		fmt.Println("超时,请重试")
		ctx.JSON(http.StatusGatewayTimeout, gin.H{
			"msg":    "超时,请重试",
			"status": 200,
		})
		return
	}
	fmt.Println("跳出 select...")
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "Success",
		"status": 200,
	})
}
func Del(ctx *gin.Context) {
	<-ch
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "Success",
		"status": 200,
	})
}
