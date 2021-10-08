package controller

import (
	"gotest/response"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func UpSingleFileHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upSingleFile.html", gin.H{"title": "test txt"})
}

func UpLoadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("f1")
	if err != nil {
		response.Fail(ctx, gin.H{"message": "上传失败"}, "")
		return
	}
	workDir, _ := os.Getwd()
	dst := workDir + "/data/" + time.Now().Format("2006-01-02-15-04-05") + "-" + file.Filename
	ctx.SaveUploadedFile(file, dst)
	response.Success(ctx, gin.H{"message": "上传成功，保存位置" + dst}, "ok")
}
