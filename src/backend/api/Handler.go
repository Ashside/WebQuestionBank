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

	// 添加题目
	var question SubjectiveQuestions
	question.Id = findAvailableQuesID(db)
	question.Content = form.Question
	question.Answer = form.Answer
	question.Difficulty = strconv.Itoa(form.Difficulty)
	question.Subject = form.Subject
	question.Author = form.Username

	if err := AddSubjectQuestion(db, &question); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"success": true, "reason": nil})

	keywords, err := getKeyword(form.Question)
	if err != nil {
		return
	}

	AddKeywords(db, keywords, question.Id, false)

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

	// 添加题目
	var question ChoiceQuestions
	question.Id = findAvailableQuesID(db)
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

	// 添加成功
	context.JSON(http.StatusOK, gin.H{"success": true, "reason": nil})

	keywords, err := getKeyword(form.Question)
	if err != nil {
		return
	}
	AddKeywords(db, keywords, question.Id, true)

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
			// 查询关键词

			keywordsS, err := GetKeywordsByQuestionId(db, question.Id, true)
			keywords := make([]map[string]string, 0)
			for _, keyword := range keywordsS {
				keywords = append(keywords, map[string]string{"keyword": keyword.Keyword})
			}
			if err != nil {
				log.Println("Failed to get keywords")
				return
			}

			diff, _ := strconv.Atoi(question.Difficulty)
			response = append(response, gin.H{
				"type":       "multipleChoice",
				"question":   question.Content,
				"answer":     question.Answer,
				"difficulty": diff,
				"subject":    question.Subject,
				"option":     option,
				"keywords":   keywords,
				"id":         question.Id,
			})
			continue
		} else {
			// 查询关键词
			keywordsS, err := GetKeywordsByQuestionId(db, question.Id, false)

			keywords := make([]map[string]string, 0)
			for _, keyword := range keywordsS {
				keywords = append(keywords, map[string]string{"keyword": keyword.Keyword})
			}
			if err != nil {
				log.Println("Failed to get keywords")
				return
			}
			diff, _ := strconv.Atoi(question.Difficulty)
			response = append(response, gin.H{
				"type":       "simpleAnswer",
				"question":   question.Content,
				"answer":     question.Answer,
				"difficulty": diff,
				"subject":    question.Subject,
				"option":     "",
				"keywords":   keywords,
				"id":         question.Id,
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
			continue
		}
		var subjectQuestion SubjectiveQuestions
		if err := db.Table("subjectivequestions").Where("id = ?", question.ID).First(&subjectQuestion).Error; err == nil {
			err := DeleteSubjectQuestion(db, question.ID)
			if err != nil {
				log.Println("Delete Subject Question Error")
				context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
				return
			}
			continue
		}
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Question not found"})
		return
	}

	log.Println("DeleteQuestionPost")

	context.JSON(http.StatusOK, gin.H{"success": true, "reason": nil})
}

func MakeTestPost(context *gin.Context) {
	log.Println("MakeTestPost")

	type Question struct {
		// 题目id
		ID int64 `json:"id"`
		// 本题对应的分值
		Score int64 `json:"score"`
	}
	type Request struct {
		// 选取的题目列表
		Questions []Question `form:"questions" binding:"required"`
		// 试卷名称
		TestName string `form:"testName" binding:"required"`
		// 提交者邮箱
		Username string `form:"username" binding:"required"`
	}
	type Response struct {
		// pdf文件的URL，后端自动生成pdf，返回一个可以访问到此pdf的URL，前端通过跳转访问此pdf文件。
		PDFURL string `json:"pdfURL"`
		// 原因，若成功为null
		Reason string `json:"reason"`
		// 是否成功
		Success bool `json:"success"`
	}

	log.Println("Binding form")
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

	// 查询用户是否存在
	var user Users
	if err := GetUserByUsername(db, request.Username, &user); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return
	}

	// 鉴权，要求学生无法创建试卷
	if user.Type == STUDENT {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Permission denied"})
		return

	}
	log.Println("Admin: ", request.Username)
	log.Println("Make Test: ", request.TestName)
	log.Println("Questions: ", request.Questions)

	var testId int
	testId = findAvailableTestsId(db)
	if testId == -1 {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error", "pdfURL": ""})
		return
	}
	// 添加试卷
	var test Tests
	test.Id = testId
	test.Name = request.TestName
	test.Author = request.Username

	for _, question := range request.Questions {
		test.Grade = float64(question.Score)
		test.QuestionId = int(question.ID)
		if err := AddTest(db, &test); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error", "pdfURL": ""})
			return
		}
	}

	// 生成pdf
	mdFile, err := GenerateMD(db, testId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error", "pdfURL": ""})
		return
	}

	pdfURL, err := GeneratePDFFile(mdFile, testId)

	context.JSON(http.StatusOK, gin.H{"success": true, "reason": nil, "pdfURL": pdfURL})

}
