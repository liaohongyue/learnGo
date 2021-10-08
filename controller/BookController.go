package controller

import (
	"gotest/response"

	"github.com/gin-gonic/gin"
)

func GetBook(ctx *gin.Context) {
	response.Success(ctx, gin.H{"message": "GET"}, "GET")
}

func PostBook(ctx *gin.Context) {
	response.Success(ctx, gin.H{"message": "POST"}, "POST")
}

func PutBook(ctx *gin.Context) {
	response.Success(ctx, gin.H{"message": "PUT"}, "PUT")
}

func DeleteBook(ctx *gin.Context) {
	response.Success(ctx, gin.H{"message": "DELETE"}, "DELETE")
}
