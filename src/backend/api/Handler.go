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
	if err := db.Table("subjectivequestions").Where("content = ?", form.Question).First(&SubjectiveQuestions{}).Error; err == nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Question already exists"})
		return
	}

	// 查询选择题和主观题的数量之和
	var cntChoice int64
	var cntSubject int64
	if err := db.Table("choicequestions").Count(&cntChoice).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}
	if err := db.Table("subjectivequestions").Count(&cntSubject).Error; err != nil {
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
	// TODO 生成题目关键词
	// TODO 添加关键词
	keywords, err := getKeyword(form.Question)
	if err != nil {
		return
	}
	AddKeywords(db, keywords, question.Id, false)
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
	if err := db.Table("choicequestions").Where("content = ?", form.Question).First(&ChoiceQuestions{}).Error; err == nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Question already exists"})
		return
	}

	// 查询选择题和主观题的数量之和
	var cntChoice int64
	var cntSubject int64
	if err := db.Table("choicequestions").Count(&cntChoice).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}
	if err := db.Table("subjectivequestions").Count(&cntSubject).Error; err != nil {
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
	// TODO 生成题目关键词
	// TODO 添加关键词
	keywords, err := getKeyword(form.Question)
	if err != nil {
		return
	}
	AddKeywords(db, keywords, question.Id, true)

	// 添加成功
	context.JSON(http.StatusOK, gin.H{"success": true, "reason": nil})

}

func QueryQuestionPost(context *gin.Context) {
	log.Println("QueryQuestionPost")
	var form struct {
		Username   string `form:"username" binding:""`
		Subject    string `form:"subject" binding:""`
		Difficulty int    `form:"difficulty" binding:""`
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

	log.Println(form.Username)
	log.Println(form.Subject)
	log.Println(form.Difficulty)
	// 查询题目
	questions := QueryQuestionFromCertainInf(db, form.Username, form.Subject, form.Difficulty)

	var response []gin.H
	for _, question := range questions {
		// 检查question的类型
		if question.Options != "" {
			// 解析question的option
			var option map[string]string
			err := json.Unmarshal([]byte(question.Options), &option)
			if err != nil {
				log.Println("JSON", question.Options)
				fmt.Println("Error unmarshalling option:", err)
				return
			}
			response = append(response, gin.H{
				"type":       "multipleChoice",
				"question":   question.Content,
				"answer":     question.Answer,
				"difficulty": question.Difficulty,
				"subject":    question.Subject,
				"option":     option,
			})
			continue
		} else {

			response = append(response, gin.H{
				"type":       "simpleAnswer",
				"question":   question.Content,
				"answer":     question.Answer,
				"difficulty": question.Difficulty,
				"subject":    question.Subject,
				"option":     "",
			})
		}

	}
	context.JSON(http.StatusOK, gin.H{"success": true, "reason": nil, "questions": response})
}

func DeleteQuestionPost(context *gin.Context) {
	type Question struct {
		// 题目id
		ID int `json:"id"`
	}
	type Request struct {
		// 选取的题目列表
		Questions []Question `form:"questions" binding:"required"`
		Username  string     `form:"username" binding:"required"`
	}
	var request Request
	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid form"})
		return
	}
	db, err := getDatabase()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}
	// 检查是否是管理员
	var user Users
	if err := GetUserByUsername(db, request.Username, &user); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return

	}
	if user.Type != ADMIN {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Permission denied"})
		return
	}
	log.Println("Admin: ", request.Username)
	log.Println("Delete Questions: ", request.Questions)

	for _, question := range request.Questions {
		// 查询题目是否存在
		var choiceQuestion ChoiceQuestions
		if err := db.Table("choicequestions").Where("id = ?", question.ID).First(&choiceQuestion).Error; err == nil {
			err := DeleteChoiceQuestion(db, question.ID)
			if err != nil {
				log.Println("Delete Choice Question Error")
				context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
				return
			}
		}
		var subjectQuestion SubjectiveQuestions
		if err := db.Table("subjectivequestions").Where("id = ?", question.ID).First(&subjectQuestion).Error; err == nil {
			err := DeleteSubjectQuestion(db, question.ID)
			if err != nil {
				log.Println("Delete Subject Question Error")
				context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
				return
			}
		}
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Question not found"})
		return
	}

	log.Println("DeleteQuestionPost")
}
