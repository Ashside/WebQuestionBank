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
	mdFile, err := GenerateMdByTestID(db, testId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error", "pdfURL": ""})
		return
	}

	pdfURL, err := GeneratePDFFile(mdFile, testId)

	context.JSON(http.StatusOK, gin.H{"success": true, "reason": nil, "pdfURL": pdfURL})

}

func QueryAllTestsPost(context *gin.Context) {
	type Request struct {
		// 用户名，要查询的用户名
		Username string `form:"username"`
	}
	type Test struct {
		// 试卷ID
		ID int64 `json:"id"`
		// 试卷名
		Name string `json:"name"`
	}
	type Response struct {
		// 原因，如果失败返回原因，如果成功则为 null
		Reason string `json:"reason"`
		// 是否成功
		Success bool   `json:"success"`
		Test    []Test `json:"test"`
	}

	log.Println("QueryAllTestsPost")
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
	var tests []Tests
	if request.Username == "" {
		var err error
		tests, err = QueryAllTests(db, "", ADMIN)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
			return

		}
	} else {
		// 查询用户是否存在
		var user Users
		if err := GetUserByUsername(db, request.Username, &user); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
			return
		}

		// 鉴权
		if user.Type == STUDENT {
			context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Permission denied"})
			return
		}

		// 查询试卷
		var err error
		tests, err = QueryAllTests(db, request.Username, user.Type)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
			return
		}
	}

	// 找出tests中互异的id
	var responseTest []Test
	var idMap = make(map[int]bool)
	for _, t := range tests {
		if _, ok := idMap[t.Id]; !ok {
			idMap[t.Id] = true
			responseTest = append(responseTest, Test{ID: int64(t.Id), Name: t.Name})
		}
	}

	var response Response
	response.Success = true
	response.Reason = ""
	response.Test = responseTest
	context.JSON(http.StatusOK, response)

}

func QueryTestByIDPost(context *gin.Context) {
	type Request struct {
		// 试卷ID，要查询的试卷ID
		TestId int64 `form:"testId" binding:"required"`
		// 用户名，要查询的用户名
		Username string `form:"username"`
	}
	type Response struct {
		// 原因，如果失败返回原因，如果成功则为 null
		Reason string `json:"reason"`
		// 是否成功
		Success bool   `json:"success"`
		Test    string `json:"test"`
	}
	log.Println("QueryTestByIDPost")
	var request Request

	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid form", "test": ""})
		return
	}

	db, err := getDatabase()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error", "test": ""})
		return
	}

	var response Response
	mdFile, err := GenerateMdByTestID(db, int(request.TestId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error", "test": ""})
		return
	}

	response.Success = true
	response.Reason = ""
	response.Test = mdFile
	context.JSON(http.StatusOK, response)

}

func GetStudentAnswersPost(context *gin.Context) {
	type Request struct {
		// 教师邮箱
		Username string `form:"username" binding:"required"`
	}
	type Response struct {
		// 标准答案
		Answer string `json:"answer"`
		// 问题
		Question string `json:"question"`
		// 问题 ID
		QuestionID int `json:"questionID"`
		// 原因，原因，如果成功则不需要填写
		Reason string `json:"reason"`
		// 本题分值
		Score int `json:"score"`
		// 学生答案
		StudentAnswer string `json:"studentAnswer"`
		// 学生邮箱
		StudentUsername string `json:"studentUsername"`
		// 是否成功
		Success bool `json:"success"`
		// 试卷 ID
		TestID int `json:"testID"`
	}

	log.Println("GetStudentAnswersPost")
	var request Request

	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid form", "test": ""})
		return
	}

	db, err := getDatabase()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error", "test": ""})
		return
	}

	// 用户是否存在
	var user Users
	if err := GetUserByUsername(db, request.Username, &user); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return
	}

	// 鉴权
	if user.Type == STUDENT {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Permission denied"})
		return
	}

	// 查询学生答案
	var assign []Assignments
	assign, _ = GetAssignsByAssignName(db, request.Username)

	var response Response
	for _, a := range assign {
		ques, bExist := QueryQuestionFromId(db, a.QuestionId)
		if !bExist {
			context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Question not found"})
			return
		}
		if ques.Options == "" {
			response.StudentAnswer = a.StuAnswer
			response.Score, _ = GetGradeByTestIdAndQuestionId(db, a.TestId, a.QuestionId)
			response.StudentUsername = a.StuName
			response.TestID = a.TestId

			response.QuestionID = a.QuestionId

			response.Question = ques.Content
			response.Answer = ques.Answer
			response.Success = true
			response.Reason = ""
			context.JSON(http.StatusOK, response)
			return
		} else {
			continue
		}
	}
	context.JSON(http.StatusOK, gin.H{"success": false, "reason": "No answer found"})

}

func SubmitScorePost(context *gin.Context) {
	type Request struct {
		// 题目 ID，提交的题目 ID
		QuestionID int `json:"questionID"`
		// 成绩
		Score int `json:"score"`
		// 学生用户名
		StudentUsername string `json:"studentUsername"`
		// 试卷 ID
		TestID int `json:"testID"`
		// 用户名，要查询的用户名
		Username string `json:"username"`
	}
	type Response struct {
		// 原因，如果失败返回原因，如果成功则为 null
		Reason string `json:"reason"`
		// 是否成功
		Success bool `json:"success"`
	}

	log.Println("SubmitScorePost")
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

	// 用户是否存在
	var user Users
	if err := GetUserByUsername(db, request.Username, &user); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return
	}

	// 鉴权
	if user.Type == STUDENT {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Permission denied"})
		return
	}

	// 提交分数
	var assign Assignments
	assign.QuestionId = request.QuestionID
	assign.Score = float64(request.Score)
	assign.StuName = request.StudentUsername
	assign.TestId = request.TestID

	if err := assign.UpdateScore(db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true, "reason": ""})

}
