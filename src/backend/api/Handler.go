package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func LoginCheckPost(c *gin.Context) {
	log.Println("LoginCheckPost")
	var form struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	log.Println("Binding form")
	if err := c.ShouldBind(&form); err != nil {
		// 打印表单的键值对
		log.Printf("form: %+v\n", form)
		// 返回错误信息
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
	log.Println("RegisterCheckPost")
	var form struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
		Type     string `form:"type" binding:"required"`
	}
	log.Println("Binding form")
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

func AddSimpleAnswerPost(context *gin.Context) {

	log.Println("AddSimpleAnswerPost")
	var form struct {
		Question   string `form:"question" binding:"required"`
		Answer     string `form:"answer" binding:"required"`
		Difficulty int    `form:"difficulty" binding:"required"`
		Subject    string `form:"subject" binding:"required"`
		Username   string `form:"username" binding:"required"`
	}
	log.Println("Binding form")
	if err := context.ShouldBind(&form); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid form"})
		return
	}

	db, err := getDatabase()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}

	// 查询用户是否存在
	var user Users
	if err := GetUserByUsername(db, form.Username, &user); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return
	}

	//TODO 查询题目是否重复
	if err := db.Table("SubjectiveQuestions").Where("content = ?", form.Question).First(&SubjectiveQuestions{}).Error; err == nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Question already exists"})
		return
	}

	// TODO 生成题目关键词
	// TODO 添加关键词

	// 查询选择题和主观题的数量之和
	var cntChoice int64
	var cntSubject int64
	if err := db.Table("ChoiceQuestions").Count(&cntChoice).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}
	if err := db.Table("SubjectiveQuestions").Count(&cntSubject).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}
	count := cntChoice + cntSubject

	// 添加题目
	var question SubjectiveQuestions
	question.Id = int(count) + 1
	question.Content = form.Question
	question.Answer = form.Answer
	question.Difficulty = strconv.Itoa(form.Difficulty)
	question.Subject = form.Subject
	question.Author = form.Username

	if err := AddSubjectQuestion(db, &question); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}

	// 添加成功
	context.JSON(http.StatusOK, gin.H{"success": true, "reason": nil})
}

func AddChoiceAnswerPost(context *gin.Context) {
	log.Println("AddChoiceAnswerPost")
	var form struct {
		Question string            `form:"question" binding:"required"`
		Answer   string            `form:"answer" binding:"required"`
		Option   map[string]string `form:"option" binding:"required"`

		Difficulty int    `form:"difficulty" binding:"required"`
		Subject    string `form:"subject" binding:"required"`
		Username   string `form:"username" binding:"required"`
	}
	log.Println("Binding form")
	if err := context.ShouldBind(&form); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid form"})
		return
	}

	db, err := getDatabase()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}

	// 查询用户是否存在
	var user Users
	if err := GetUserByUsername(db, form.Username, &user); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return
	}

	//TODO 查询题目是否重复
	if err := db.Table("ChoiceQuestions").Where("content = ?", form.Question).First(&ChoiceQuestions{}).Error; err == nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Question already exists"})
		return
	}

	// TODO 生成题目关键词
	// TODO 添加关键词

	// 查询选择题和主观题的数量之和
	var cntChoice int64
	var cntSubject int64
	if err := db.Table("ChoiceQuestions").Count(&cntChoice).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}
	if err := db.Table("SubjectiveQuestions").Count(&cntSubject).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}
	count := cntChoice + cntSubject

	// 添加题目
	var question ChoiceQuestions
	question.Id = int(count) + 1
	question.Content = form.Question
	question.Answer = form.Answer
	question.Difficulty = strconv.Itoa(form.Difficulty)
	question.Subject = form.Subject
	question.Author = form.Username

	// 使用json.Marshal将其转换为JSON格式的字符串
	optionBytes, err := json.Marshal(form.Option)
	if err != nil {
		fmt.Println("Error marshalling option:", err)
		return
	}

	// 将[]byte转换为string
	optionString := string(optionBytes)
	question.Options = optionString

	if err := AddChoiceQuestion(db, &question); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return

	}
}
