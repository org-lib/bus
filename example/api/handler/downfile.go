package handler

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func DownLoadFile(ctx *gin.Context) {
	_, fname := filepath.Split("/Users/yuandeqiao/Desktop/Relese.pptx")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+fname)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.File("/Users/yuandeqiao/Desktop/Relese.pptx")
}
