package main

import (
	"github.com/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
func main() {

	r := gin.Default()
	r.Use(corsMiddleware())

	apiGroup := r.Group("/api")

	// 处理/api/usr路由组
	usrGroup := apiGroup.Group("/usr")
	// 处理/api/usr/loginCheck的OPTIONS请求
	// 该请求用于检查用户登录状态
	// 输入：form表单，包含username和password
	// 输出：json格式，包含success、reason和type字段
	//usrGroup.OPTIONS("/loginCheck", api.LoginCheckPost)
	usrGroup.POST("/loginCheck", api.LoginCheckPost)
	// 处理/api/usr/registerCheck的OPTIONS请求
	// 该请求用于检查用户注册状态
	// 输入：form表单，包含username, password, type字段
	// 输出：json格式，包含success、reason字段
	usrGroup.POST("/registerCheck", api.RegisterCheckPost)

	// 处理/api/questionBank路由组
	questionBankGroup := apiGroup.Group("/questionBank")
	// 处理/api/questionBank/addQuestion路由组
	addQuestionGroup := questionBankGroup.Group("/addQuestion")
	// 处理/api/questionBank/addQuestion/simpleAnswer的OPTIONS请求
	// 该请求用于添加题目
	// 输入：form表单，包含question, answer, difficulty, subject, username字段
	// 输出：json格式，包含success、reason字段
	addQuestionGroup.POST("/simpleAnswer", api.AddSimpleAnswerPost)
	// 处理/api/questionBank/addQuestion/choiceAnswer的OPTIONS请求
	// 该请求用于添加题目
	// 输入：form表单，包含question, answer, option, difficulty, subject, username字段
	// 输出：json格式，包含success、reason字段
	addQuestionGroup.POST("/multipleChoice", api.AddChoiceAnswerPost)

	// 处理/api/questionBank/queryQuestion路由组
	queryQuestionGroup := questionBankGroup.Group("/queryQuestion")
	// 处理/api/questionBank/queryQuestion的OPTIONS请求
	// 该请求用于查询题目
	// 输入：form表单，包含username, subject, difficulty字段.
	// 输出：json格式，包含success、reason、questions字段，questions字段是一个数组，包含多个题目
	queryQuestionGroup.POST("", api.QueryQuestionPost)

	//处理/api/questionBank/deleteQuestion的OPTIONS请求
	// 该请求用于删除题目
	// 输入：form表单，包含questions字段，questions是一个数组，包含多个题目的id
	// 输出：json格式，包含success、reason字段
	// 约束：只有管理员可以删除题目
	questionBankGroup.POST("/deleteQuestion", api.DeleteQuestionPost)

	//处理/api/questionBank/makeTest的OPTIONS请求
	// 该请求用于组卷
	// 输入：form表单，包含questions字段，questions是一个数组，包含多个题目的id, username, testName字段
	// 输出：json格式，包含success、reason、pdfURL字段
	// 约束：只有教师可以组卷
	questionBankGroup.POST("/makeTest", api.MakeTestPost)

	//处理/api/questionBank/queryAllTests的OPTIONS请求
	// 该请求用于查询所有试卷
	// 输入：form表单，包含username字段
	// 输出：json格式，包含success、reason、test字段，test字段是一个数组，包含多个试卷
	questionBankGroup.POST("/queryAllTests", api.QueryAllTestsPost)

	//处理/api/questionBank/queryTestByID的OPTIONS请求
	// 该请求用于查询试卷下所有题目
	// 输入：form表单，包含username, id字段，id为试卷编号
	// 输出：json格式，包含success、reason、test字段，test字段是一个字符串
	questionBankGroup.POST("/queryTestByID", api.QueryTestByIDPost)
	//处理/api/questionBank/findSameTestByID的OPTIONS请求
	// 该请求用于查找相似试卷
	// 输入：form表单，包含username,id字段，id为试卷编号
	// 输出：json格式，包含success、reason、questions字段，questions字段是一个数组，包含多个题目
	questionBankGroup.POST("/findSameTestByID", api.FindSamePost)

	_ = r.Run(":8081")
}
