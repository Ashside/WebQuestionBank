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
		Name     string `form:"name" binding:"required"`
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
	user.Name = form.Name
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

	keywords, err := getKeywordFromLocal(form.Question)
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

	keywords, err := getKeywordFromLocal(form.Question)
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
		QuestionID int64 `json:"questionID"`
		// 原因，原因，如果成功则不需要填写
		Reason string `json:"reason"`
		// 本题分值
		Score int64 `json:"score"`
		// 学生答案
		StudentAnswer string `json:"studentAnswer"`
		// 学生邮箱
		StudentUsername string `json:"studentUsername"`
		// 是否成功
		Success bool `json:"success"`
		// 试卷 ID
		TestID int64 `json:"testID"`
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
	//log.Println(assign)
	var response Response
	for _, a := range assign {
		if a.StuScore != -1 {
			continue
		}
		ques, bExist := QueryQuestionFromId(db, a.QuestionId)
		if !bExist {
			context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Question not found"})
			return
		}
		log.Println(ques)
		if ques.Options == "" {

			response = Response{
				Success:         true,
				Reason:          "",
				Answer:          ques.Answer,
				Question:        ques.Content,
				QuestionID:      int64(ques.Id),
				Score:           int64(a.Score),
				StudentAnswer:   a.StuAnswer,
				StudentUsername: a.StuName,
				TestID:          int64(a.TestId),
			}
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
		QuestionID int64 `form:"questionID" binding:"required"`
		// 成绩
		Score int64 `form:"score" binding:"required"`
		// 学生用户名
		StudentUsername string `form:"studentUsername" binding:"required"`
		// 试卷 ID
		TestID int64 `form:"testID" binding:"required"`
		// 用户名，要查询的用户名
		Username string `form:"username" binding:"required"`
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

	// 验证题目是否存在
	_, bExist := QueryQuestionFromId(db, int(request.QuestionID))
	if !bExist {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Question not found"})
		return

	}

	// 验证分数是否合法
	quesScore, _ := QueryGradeByTestIdAndQuestionId(db, int(request.TestID), int(request.QuestionID))
	if request.Score < 0 || request.Score > 100 || request.Score > int64(quesScore) {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Invalid score"})
		return

	}

	// 提交分数
	var assign Assignments
	assign.QuestionId = int(request.QuestionID)
	assign.StuName = request.StudentUsername
	assign.TestId = int(request.TestID)
	assign.StuScore = float64(request.Score)
	if err := assign.UpdateScore(db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true, "reason": ""})

}

func DistributeTestPost(context *gin.Context) {
	type Request struct {
		// 学生 ID 数组
		Students []string `form:"students" binding:"required"`
		// 测试名
		TestID int `form:"testID" binding:"required"`
		// 提交教师名
		Username string `form:"username" binding:"required"`
	}
	type Response struct {
		// 原因，成功原因为空串
		Reason string `json:"reason"`
		// 是否成功
		Success bool `json:"success"`
	}

	log.Println("DistributeTestPost")
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

	// 分发试卷
	for _, student := range request.Students {
		var assign Assignments
		quesId := QueryQuesIdByTestID(db, request.TestID)
		for _, q := range quesId {
			assign.QuestionId = q
			assign.StuName = student
			assign.TestId = request.TestID
			quesScore, _ := QueryGradeByTestIdAndQuestionId(db, request.TestID, q)
			assign.Score = quesScore
			assign.StuAnswer = ""
			assign.StuScore = -1
			assign.AssignName = request.Username
			assign.Finished = false
			if err := assign.AddAssign(db); err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Error in distributing test"})
				return
			}

		}

	}

	context.JSON(http.StatusOK, gin.H{"success": true, "reason": ""})

}

func FindAllStudentsPost(context *gin.Context) {
	type Request struct {
		// 老师 username
		Username string `form:"username" binding:"required"`
	}
	type Student struct {
		// 学生姓名
		Student string `json:"student"`
		// 学生邮箱
		StudentUsername string `json:"studentUsername"`
	}
	type Response struct {
		// 原因，如果成功原因为空
		Reason string `json:"reason"`
		// 学生数组
		Students []Student `json:"students"`
		// 是否成功
		Success bool `json:"success"`
	}

	log.Println("FindAllStudentsPost")
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

	// 查询学生
	var students []Users
	students, err = QueryAllStudents(db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}

	var response Response
	response.Success = true
	response.Reason = ""
	for _, s := range students {
		response.Students = append(response.Students, Student{Student: s.Name, StudentUsername: s.Username})

	}
	context.JSON(http.StatusOK, response)

}

func QueryAllTestsByStudentIDPost(context *gin.Context) {
	type Request struct {
		// 用户名，要查询的学生用户名
		Username string `form:"username" binding:"required"`
	}
	type Test struct {
		// 试卷ID
		ID int64 `json:"id"`
		// 试卷名
		Name string `json:"name"`
		// 当前试卷状态，两个状态，to_be_finish 和 finish
		State string `json:"state"`
	}
	type Response struct {
		// 原因，如果失败返回原因，如果成功则为 null
		Reason string `json:"reason"`
		// 是否成功
		Success bool   `json:"success"`
		Test    []Test `json:"test"`
	}

	log.Println("QueryAllTestsByStudentIDPost")
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

	// 查询试卷
	var assign []Assignments
	assign, _ = GetAssignsByStuName(db, request.Username)

	// 获得互异的试卷ID
	var testIdMap = make(map[int]bool)
	var testId []int
	for _, a := range assign {
		if _, ok := testIdMap[a.TestId]; !ok {
			testIdMap[a.TestId] = true
			testId = append(testId, a.TestId)
		}

	}

	var response Response
	for _, t := range testId {
		if isTestFinished(db, t, request.Username) {
			response.Test = append(response.Test, Test{ID: int64(t), Name: QueryTestNameByTestId(db, t), State: "finish"})
		} else {
			response.Test = append(response.Test, Test{ID: int64(t), Name: QueryTestNameByTestId(db, t), State: "to_be_finish"})
		}
	}

	response.Success = true
	response.Reason = ""
	context.JSON(http.StatusOK, response)

}

func QueryTestStateByStudentIDPost(context *gin.Context) {
	type Request struct {
		// 要查的学生 ID
		StudentUsername string `form:"studentUsername"`
		// 要查的测试 ID
		TestID string `form:"testID"`
	}
	// 选择题选项，选择题需要填写，简答题不需要
	type Option struct {
		Option1 string `json:"option1"`
		Option2 string `json:"option2"`
		Option3 string `json:"option3"`
		Option4 string `json:"option4"`
	}
	type Question struct {
		// 答案
		Answer string `json:"answer"`
		// 题目id
		ID int64 `json:"id"`
		// 选择题选项，选择题需要填写，简答题不需要
		Option *Option `json:"option,omitempty"`
		// 问题内容
		Question string `json:"question"`
		// 学生存储中的答案
		StudentAnswer string `json:"studentAnswer"`
		// 问题类型，简答题和选择题两种
		Type string `json:"type"`
	}
	type Response struct {
		// 问题列表
		Questions []Question `json:"questions"`
		// 原因，如果失败返回原因，如果成功则为 null
		Reason string `json:"reason"`
		// 是否成功
		Success bool `json:"success"`
	}

	log.Println("QueryTestStateByStudentIDPost")
	var request Request
	if err := context.ShouldBind(&request); err != nil {
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
	if err := GetUserByUsername(db, request.StudentUsername, &user); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return
	}

	// 查询试卷
	var assign []Assignments
	tId, _ := strconv.Atoi(request.TestID)
	assign, _ = QueryAssignsByTestAndStu(db, tId, request.StudentUsername)

	var response Response
	for _, a := range assign {
		ques, bExist := QueryQuestionFromId(db, a.QuestionId)
		if !bExist {
			context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Question not found"})
			return
		}
		if ques.Options == "" {
			response.Questions = append(response.Questions, Question{
				Answer:        ques.Answer,
				ID:            int64(ques.Id),
				Option:        nil,
				Question:      ques.Content,
				StudentAnswer: a.StuAnswer,
				Type:          "simpleAnswer",
			})
		} else {
			var option Option
			err := json.Unmarshal([]byte(ques.Options), &option)
			if err != nil {
				fmt.Println("Error unmarshalling option:", err)
				return
			}
			response.Questions = append(response.Questions, Question{
				Answer: ques.Answer,
				ID:     int64(ques.Id),
				Option: &option,

				Question:      ques.Content,
				StudentAnswer: a.StuAnswer,
				Type:          "multipleChoice",
			})
		}

	}

	response.Success = true
	response.Reason = ""
	context.JSON(http.StatusOK, response)

}

func SaveTestAnswerByStudentIDPost(context *gin.Context) {
	type Question struct {
		// 题目ID
		ID int64 `json:"id"`
		// 学生答案
		StudentAnswer string `json:"studentAnswer"`
		// 题目类型
		Type string `json:"type"`
	}
	type Request struct {
		Questions []Question `form:"questions"`
		// 学生用户名
		StudentUsername string `form:"studentUsername"`
		// 测试ID
		TestID int64 `form:"testID"`
	}

	type Response struct {
		// 原因，如果失败返回原因，如果成功则为 null
		Reason string `json:"reason"`
		// 是否成功
		Success bool `json:"success"`
	}

	log.Println("SaveTestAnswerByStudentIDPost")
	var request Request
	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid form"})
		return
	}

	// 查询试卷
	db, err := getDatabase()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}

	// 查询用户是否存在
	var user Users
	if err := GetUserByUsername(db, request.StudentUsername, &user); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return
	}

	// 更新学生答案
	for _, q := range request.Questions {
		var assign Assignments
		assign.QuestionId = int(q.ID)
		assign.StuName = request.StudentUsername
		assign.TestId = int(request.TestID)
		assign.StuAnswer = q.StudentAnswer
		assign.StuScore = CheckScore(db, assign)
		log.Println(assign.StuScore)
		assign.Finished = false
		if err := assign.UpdateAnswer(db); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
			return
		}
		if err := assign.UpdateScore(db); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
			return
		}
		if err := assign.UpdateFinished(db); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
			return
		}

	}

	context.JSON(http.StatusOK, gin.H{"success": true, "reason": ""})
}

func SubmitTestAnswerByStudentIDPost(context *gin.Context) {

	type Question struct {
		// 题目ID
		ID int64 `json:"id"`
		// 学生答案
		StudentAnswer string `json:"studentAnswer"`
		// 题目类型
		Type string `json:"type"`
	}
	type Request struct {
		Questions []Question `form:"questions"`
		// 学生用户名
		StudentUsername string `form:"studentUsername"`
		// 测试ID
		TestID int64 `form:"testID"`
	}

	type Response struct {
		// 原因，如果失败返回原因，如果成功则为 null
		Reason string `json:"reason"`
		// 是否成功
		Success bool `json:"success"`
	}

	log.Println("SubmitTestAnswerByStudentIDPost")
	var request Request
	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid form"})
		return
	}

	// 查询试卷
	db, err := getDatabase()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
		return
	}

	// 查询用户是否存在
	var user Users
	if err := GetUserByUsername(db, request.StudentUsername, &user); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return
	}

	// 更新学生答案
	for _, q := range request.Questions {
		var assign Assignments
		assign.QuestionId = int(q.ID)
		assign.StuName = request.StudentUsername
		assign.TestId = int(request.TestID)
		assign.StuAnswer = q.StudentAnswer
		assign.Finished = true
		assign.StuScore = CheckScore(db, assign)

		if err := assign.UpdateAnswer(db); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
			return
		}
		if err := assign.UpdateScore(db); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
			return
		}
		if err := assign.UpdateFinished(db); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Internal error"})
			return
		}

	}

	context.JSON(http.StatusOK, gin.H{"success": true, "reason": ""})
}

func QueryTestDetailByStudentIDPost(context *gin.Context) {
	type Request struct {
		// 学生用户名
		StudentUsername string `form:"studentUsername"`
		// 测试 ID
		TestID string `form:"testID"`
	}
	// 选项
	type Option struct {
		Option1 string `json:"option1"`
		Option2 string `json:"option2"`
		Option3 string `json:"option3"`
		Option4 string `json:"option4"`
	}
	type Question struct {
		// 问题标准答案
		Answer string `json:"answer"`
		// 满分
		FullScore int64 `json:"fullScore"`
		// 问题 ID
		ID int64 `json:"id"`
		// 判卷是否完成
		IsReviewComplete bool `json:"isReviewComplete"`
		// 选项
		Option *Option `json:"option,omitempty"`
		// 问题题干
		Question string `json:"question"`
		// 学生答案
		StudentAnswer string `json:"studentAnswer"`
		// 学生得分
		StudentScore int64 `json:"studentScore"`
		// 问题类型
		Type string `json:"type"`
	}
	type Response struct {
		// 问题列表
		Questions []Question `json:"questions"`
		// 原因
		Reason string `json:"reason"`
		// 是否成功
		Success bool `json:"success"`
	}

	log.Println("QueryTestDetailByStudentIDPost")
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
	if err := GetUserByUsername(db, request.StudentUsername, &user); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Users not found"})
		return
	}

	// 查询试卷
	var assign []Assignments
	tID, _ := strconv.Atoi(request.TestID)
	assign, _ = QueryAssignsByTestAndStu(db, tID, request.StudentUsername)

	var response Response

	for _, a := range assign {
		ques, bExist := QueryQuestionFromId(db, a.QuestionId)
		if !bExist {
			context.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "Question not found"})
			return
		}
		if ques.Options == "" {
			response.Questions = append(response.Questions, Question{
				Answer:           ques.Answer,
				FullScore:        int64(a.Score),
				ID:               int64(ques.Id),
				IsReviewComplete: a.StuScore != -1,
				Option:           nil,
				Question:         ques.Content,
				StudentAnswer:    a.StuAnswer,
				StudentScore:     int64(a.StuScore),
				Type:             "simpleAnswer",
			})
		} else {
			var option Option
			err := json.Unmarshal([]byte(ques.Options), &option)
			if err != nil {
				fmt.Println("Error unmarshalling option:", err)
				return
			}
			response.Questions = append(response.Questions, Question{
				Answer:           ques.Answer,
				FullScore:        int64(a.Score),
				ID:               int64(ques.Id),
				IsReviewComplete: a.StuScore != -1,
				Option:           &option,
				Question:         ques.Content,
				StudentAnswer:    a.StuAnswer,
				StudentScore:     int64(a.StuScore),
				Type:             "multipleChoice",
			})
		}
	}

	response.Success = true
	response.Reason = ""
	context.JSON(http.StatusOK, response)

}
