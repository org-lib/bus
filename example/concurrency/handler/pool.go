package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/org-lib/bus/config"
	"net/http"
)

func AddPool(ctx *gin.Context) {
	config.Work.Add(1)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "Success",
		"status": 200,
	})
}
func DelPool(ctx *gin.Context) {
	config.Work.Done()
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "Success",
		"status": 200,
	})
}
