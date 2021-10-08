package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"gotest/common"
)

func main() {
	InitConfig()
	common.InitDB()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.LoadHTMLFiles("templates/upSingleFile.html")
	r = CollectRouters(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
