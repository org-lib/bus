package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "Success",
		"status": 200,
	})
}
