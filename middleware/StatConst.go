package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func StatusConst() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Set("name", "xiaowangzi")
		ctx.Next()
		cost := time.Since(start)
		fmt.Println(cost)
	}
}
