package api

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"log"
)
type QuestionSummary struct {
	ID         int   `json:"id"`
	QuestionType string `json:"question_type"`
}
// 从题目列表中随机选择指定数量的题目
func SelectRandomQuestions(questions []QuestionSummary, count int) []QuestionSummary {
	if len(questions) <= count {
		return questions
	}
	rand.Seed(time.Now().UnixNano())

	for i := range questions {
		j := rand.Intn(i + 1)
		questions[i], questions[j] = questions[j], questions[i]
	}

	return questions[:count]
}
// 根据不同的条件查询题目
func QueryQuestions(db *gorm.DB, tableName, selectFields string, conditions []string, args ...interface{}) ([]QuestionSummary, error) {
	log.Printf("Querying %s with conditions: %v", tableName, conditions)

	var questions []QuestionSummary
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	joinClause := fmt.Sprintf(`
		JOIN %s_keywords ON %s_keywords.question_id = %s.id
		JOIN keywords ON keywords.id = %s_keywords.keyword_id
	`, tableName, tableName, tableName, tableName) 

	err := db.Table(tableName).
		Select(selectFields).
		Joins(joinClause).
		Where(whereClause, args...).
		Group(tableName + ".id"). 
		Scan(&questions).Error

	if err != nil {
		log.Printf("Error querying %s: %v", tableName, err)
	} else {
		log.Printf("Successfully queried %d %s questions", len(questions), tableName)
	}

	return questions, err
}
func SearchQuestions(c *gin.Context) {
	var form struct {
		Username   string `form:"username" binding:"required"`
		Difficulty int `form:"difficulty,omitempty"` // 可选参数，使用omitempty避免未提供时产生错误
		Subject    string `form:"subject,omitempty"`    // 可选参数
		Keyword    string `form:"keyword,omitempty"`    // 可选参数
	}

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid form data"})
		return
	}

	db, err := getDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to connect to database"})
		return
	}

	// 确保用户存在
	var user Users
	if err := GetUserByUsername(db, form.Username, &user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "reason": "User not found"})
		return
	}

	// 构建查询条件
	var conditions1 []string
	var conditions2 []string
	args := make([]interface{}, 0)
	
	if form.Difficulty != 0 { 
		conditions1 = append(conditions1, "choice_questions.difficulty = ?")
		conditions2 = append(conditions2, "subjective_questions.difficulty = ?")
		args = append(args, form.Difficulty)
	}
	if form.Subject != "" {
		conditions1 = append(conditions1, "choice_questions.subject = ?")
		conditions2 = append(conditions2, "subjective_questions.subject = ?")
		args = append(args, form.Subject)
	}
	if form.Keyword != "" {
		conditions1 = append(conditions1, "keywords.keyword = ?")
		conditions2 = append(conditions2, "keywords.keyword = ?")
		args = append(args, form.Keyword)
	}

	log.Printf("Searching questions for user: %s with criteria: %+v", form.Username, form)

	// 查询选择题
	choiceQuestions, err := QueryQuestions(db, "choice_questions", "choice_questions.id AS id, 'choice' AS question_type", conditions1, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to fetch choice questions by criteria"})
		return
	}

	// 查询主观题
	subjectiveQuestions, err := QueryQuestions(db, "subjective_questions", "subjective_questions.id AS id, 'subjective' AS question_type", conditions2, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to fetch subjective questions by criteria"})
		return
	}

	choiceQuestions = SelectRandomQuestions(choiceQuestions, 10)
	subjectiveQuestions = SelectRandomQuestions(subjectiveQuestions, 10)
	questions := append(choiceQuestions, subjectiveQuestions...)

	// 返回结果
	if len(questions) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "No questions matched the given criteria.",
			"questions": []QuestionSummary{},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "questions": questions})
	
	log.Println("Question search completed successfully")
}

