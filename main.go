package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Password  string `gorm:"size:255;not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
}

func main() {
	DB := InitDB()
	defer DB.Close()

	r := gin.Default()

	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		password := ctx.PostForm("password")
		telephone := ctx.PostForm("telephone")

		// 验证数据
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能小于6位"})
			return
		}
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须11位"})
			return
		}
		if len(name) == 0 {
			name = randString(10)
		}

		if isTelephoneExist(DB, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}

		fmt.Println(name, password, telephone)

		//创建用户
		newUser := User{
			Name:      name,
			Password:  password,
			Telephone: telephone,
		}
		DB.Create(&newUser)
		ctx.JSON(200, gin.H{"code": 200, "msg": "注册成功"})
	})

	panic(r.Run())
}

func randString(n int) string {
	rand.Seed(time.Now().Unix())
	letters := []byte("asdfghjklqwertyuiopzxcvbnm")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "test"
	user := "root"
	password := ""
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True", user, password, host, port, database, charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed connect database" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone=?", telephone).First(&user)
	return user.ID != 0
}
