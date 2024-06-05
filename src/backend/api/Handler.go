package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginCheckPost(c *gin.Context) {

	var form struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid form", "type": "null"})
		return
	}

	db, err := getDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error", "type": "null"})
		return

	}
	// 查询用户是否存在
	var user Users
	if err := GetUserByUsername(db, form.Username, &user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found", "type": "null"})
		return
	}

	// 检查密码是否正确
	if form.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Incorrect password", "type": "null"})
		return
	}

	// 登录成功
	c.JSON(http.StatusOK, gin.H{"success": true, "reason": nil, "type": user.Type})
}

func RegisterCheckPost(c *gin.Context) {

	var form struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
		Type     string `form:"type" binding:"required"`
	}
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid form"})
		return
	}

	db, err := getDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}

	// 查询用户是否存在
	var user Users
	if err := GetUserByUsername(db, form.Username, &user); err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users already exists"})
		return
	}

	// 注册用户
	user.Username = form.Username
	user.Password = form.Password
	user.Type = form.Type
	if err := AddUser(db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}

	// 注册成功
	c.JSON(http.StatusOK, gin.H{"success": true, "reason": nil})
}
