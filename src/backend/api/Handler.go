package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginCheckPost(context *gin.Context) {

	types := context.DefaultPostForm("type", "post")
	email := context.PostForm("email")
	password := context.PostForm("password")

	context.String(http.StatusOK, fmt.Sprintf("username:%s , password:%s , types:%s", email, password, types))

}
