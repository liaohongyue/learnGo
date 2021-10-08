package controller

import (
	"fmt"
	"gotest/common"
	"gotest/dto"
	"gotest/model"
	"gotest/response"
	"gotest/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {

	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	telephone := ctx.PostForm("telephone")

	// 验证数据
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须11位")
		return
	}
	if len(name) == 0 {
		name = utils.RandString(10)
	}

	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	fmt.Println(name, password, telephone)

	//创建用户
	hasedPasword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	newUser := model.User{
		Name:      name,
		Password:  string(hasedPasword),
		Telephone: telephone,
	}
	DB.Create(&newUser)
	response.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) {
	// 获取参数
	DB := common.GetDB()
	// 获取参数
	password := ctx.PostForm("password")
	telephone := ctx.PostForm("telephone")

	// 验证数据
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须11位")
		return
	}

	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// 验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		fmt.Println(err.Error())
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		return
	}

	// 返回数据
	response.Success(ctx, gin.H{"token": token}, "注册成功")

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	return user.ID != 0
}

func QueryUser(ctx *gin.Context) {
	username := ctx.DefaultQuery("username", "zhangsan")
	address := ctx.Query("address")
	response.Success(ctx, gin.H{"username": username, "address": address}, "ok")
}

func SearchUser(ctx *gin.Context) {
	username := ctx.PostForm("username")
	address := ctx.PostForm("address")
	response.Success(ctx, gin.H{"username": username, "address": address}, "ok")
}

func GetUser(ctx *gin.Context) {
	username := ctx.Param("username")
	address := ctx.Param("address")
	response.Success(ctx, gin.H{"username": username, "address": address}, "ok")
}

type UserInfo struct {
	Username string `form:"username" binding:"required"`
	Address  string `form:"address" binding:"required"`
}

func AllgetUser(ctx *gin.Context) {
	var userInfo UserInfo
	if err := ctx.ShouldBind(&userInfo); err == nil {
		response.Success(ctx, gin.H{"username": userInfo.Username, "address": userInfo.Address}, "ok")
	}
}

func AllPostUser(ctx *gin.Context) {
	var userInfo UserInfo
	if err := ctx.ShouldBind(&userInfo); err == nil {
		response.Success(ctx, gin.H{"username": userInfo.Username, "address": userInfo.Address}, "ok")
	}
}
