package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/org-lib/bus/logger"
	"net/http"
)

func HelloWorldPost(ctx *gin.Context) {
	var tmpMsgsUser MsgsUser
	if err := ctx.ShouldBindJSON(&tmpMsgsUser); err != nil {
		logger.Log.Error(fmt.Sprintf("DB OpenAI ： 参数解析失败：%v", err.Error()))
		ctx.JSON(http.StatusOK, MyResponse{
			Code: -1,
			Msg:  fmt.Sprintf("DB OpenAI ： 参数解析失败：%v", err.Error()),
			Data: RespBody{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "Success",
		"status": 200,
	})
}

type RequestBody struct {
	Messages         []MsgsUser `json:"messages" gorm:"column:messages"`
	MaxTokens        int        `json:"max_tokens" gorm:"column:max_tokens"`
	Temperature      float32    `json:"temperature" gorm:"column:temperature"`
	FrequencyPenalty int        `json:"frequency_penalty" gorm:"column:frequency_penalty"`
	PresencePenalty  int        `json:"presence_penalty" gorm:"column:presence_penalty"`
	TopP             float32    `json:"top_p" gorm:"column:top_p"`
	Stop             []string   `json:"stop" gorm:"column:stop"`
}
type MsgsUser struct {
	Role    string `json:"role" gorm:"column:role" length:"4"`
	Content string `json:"content" gorm:"column:content"`
}

// 响应
type RespBody struct {
	Id      string   `json:"id" gorm:"column:id"`
	Object  string   `json:"object" gorm:"column:object"`
	Created int64    `json:"created" gorm:"column:created"`
	Model   string   `json:"model" gorm:"column:model"`
	Choices []Choice `json:"choices" gorm:"column:choices"`
	UsAge   Usage    `json:"usage" gorm:"column:usage"`
}

type Choice struct {
	Index        int      `json:"index" gorm:"column:database"`
	FinishReason string   `json:"finish_reason" gorm:"column:database"`
	Msgs         MsgsUser `json:"message" gorm:"column:database"`
}
type Usage struct {
	CompletionTokens int `json:"completion_tokens" gorm:"column:completion_tokens"`
	PromptTokens     int `json:"prompt_tokens" gorm:"column:prompt_tokens"`
	TotalTokens      int `json:"total_tokens" gorm:"column:total_tokens"`
}

// myapi
type MyRequest struct {
}

type MyResponse struct {
	Code int      `json:"statu" gorm:"column:statu"` //0:接收成功，-1：接收失败
	Msg  string   `json:"msg" gorm:"column:msg"`
	Data RespBody `json:"data" gorm:"column:data"` //openai 返回的内容
}
