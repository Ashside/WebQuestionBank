package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 从题目列表中随机选择指定数量的题目
func SelectRandomQuestions(questions []ChoiceQuestions, count int) []ChoiceQuestions {
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
func QueryQuestions(db *gorm.DB, tableName, tableName2, selectFields string, conditions []string, args ...interface{}) ([]ChoiceQuestions, error) {
	log.Printf("Querying %s with conditions: %v", tableName, conditions)

	var questions []ChoiceQuestions
	whereClause := ""
	if len(conditions) > 0 {
		whereClause = strings.Join(conditions, " AND ")
	}

	joinClause := fmt.Sprintf(`
		JOIN %s_keywords ON %s_keywords.question_id = %s.id
		JOIN keywords ON keywords.id = %s_keywords.keyword_id
	`, tableName2, tableName2, tableName, tableName2)

	if tableName == "choicequestions" {
		selectFields = fmt.Sprintf("%s.id as id,  %s.content , %s.options, %s.subject, %s.difficulty, %s.author", tableName, tableName, tableName, tableName, tableName, tableName)
	} else {
		selectFields = fmt.Sprintf("%s.id as id,  %s.content , %s.difficulty, %s.subject, %s.author", tableName, tableName, tableName, tableName, tableName)
	}

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
		Difficulty string `form:"difficulty,omitempty"` // 可选参数，使用omitempty避免未提供时产生错误
		Subject    string `form:"subject,omitempty"`    // 可选参数
		Keyword    string `form:"keyword,omitempty"`    // 可选参数
	}

	if err := c.ShouldBind(&form); err != nil {
		log.Println("Error binding form data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "reason": "Invalid form data"})
		return
	}

	db, err := getDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to connect to database"})
		return
	}

	log.Println("begin searching")
	// 构建查询条件
	var conditions1 []string
	var conditions2 []string
	args := make([]interface{}, 0)

	if form.Difficulty != "all" {
		conditions1 = append(conditions1, "choicequestions.difficulty = ?")
		conditions2 = append(conditions2, "subjectivequestions.difficulty = ?")
		args = append(args, form.Difficulty)
	}
	if form.Subject != "all" {
		conditions1 = append(conditions1, "choicequestions.subject = ?")
		conditions2 = append(conditions2, "subjectivequestions.subject = ?")
		args = append(args, form.Subject)
	}
	if form.Keyword != "" {
		conditions1 = append(conditions1, "keywords.keyword = ?")
		conditions2 = append(conditions2, "keywords.keyword = ?")
		args = append(args, form.Keyword)
	}
	log.Printf("Searching questions with criteria: %+v", form)

	// 查询选择题
	choiceQuestions, err := QueryQuestions(db, "choicequestions", "choice_question", "*", conditions1, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to fetch choice questions by criteria"})
		return
	}

	// 查询主观题
	subjectiveQuestions, err := QueryQuestions(db, "subjectivequestions", "subjective_question", "*", conditions2, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "reason": "Failed to fetch subjective questions by criteria"})
		return
	}

	choiceQuestions = SelectRandomQuestions(choiceQuestions, 10)
	subjectiveQuestions = SelectRandomQuestions(subjectiveQuestions, 10)
	questions := append(choiceQuestions, subjectiveQuestions...)

	var response []gin.H
	// 查询题目
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
	c.JSON(http.StatusOK, gin.H{"success": true, "reason": nil, "questions": response})

	log.Println("Question search completed successfully")
}
