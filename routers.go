package main

import (
	"gotest/controller"
	"gotest/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouters(r *gin.Engine) *gin.Engine {
	auths := r.Group("/api/auth")
	{
		auths.POST("/register", controller.Register)
		auths.POST("/login", controller.Login)
		auths.GET("/info", middleware.AuthMiddleware(), middleware.StatusConst(), controller.Info)
	}

	r.GET("/api/book", controller.GetBook)
	r.POST("/api/book", controller.PostBook)
	r.PUT("/api/book", controller.PutBook)
	r.DELETE("/api/book", controller.DeleteBook)

	r.GET("/api/user/search", controller.QueryUser)
	r.POST("/api/user/searchUser", controller.SearchUser)
	r.GET("/api/user/getUser/:username/:address", controller.GetUser)
	r.GET("/api/user/allGetUser", controller.AllgetUser)
	r.POST("/api/user/allPostUser", controller.AllPostUser)

	r.GET("/upSingleFileForm", controller.UpSingleFileHTML)
	r.POST("/upLoadFile", controller.UpLoadFile)

	r.GET("/baidu", controller.Tobaidu)
	return r
}
