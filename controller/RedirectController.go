package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Tobaidu(ctx *gin.Context) {
	ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}
