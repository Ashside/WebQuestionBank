package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoginCheckPost(context *gin.Context) {

	// 获取POST请求参数

	// 获取POST请求参数中的username
	email := context.PostForm("email")
	password := context.PostForm("password")

	fmt.Println("email: ", email)
	fmt.Println("password: ", password)

}
