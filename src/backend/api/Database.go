package api

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	DatabaseUserName = "root"
	DatabasePassword = "Aa=12345678"
	//DatabasePassword = "123456"
	DatabaseName = "SEProject"

	DatabaseAddress = "121.43.124.218:3306"
	//DatabaseAddress = "localhost:3306"
)

type Users struct {
	Username string `gorm:"primaryKey"`
	Password string
	Type     string
}
type Keywords struct {
	Id      int `gorm:"primaryKey"`
	Keyword string
}
type ChoiceQuestions struct {
	Id         int `gorm:"primaryKey"`
	Subject    string
	Content    string
	Options    string
	Answer     string
	Difficulty string
	Author     string
}
type SubjectiveQuestions struct {
	Id         int `gorm:"primaryKey"`
	Subject    string
	Content    string
	Answer     string
	Difficulty string
	Author     string
}
type ChoiceQuestionKeywords struct {
	QuestionId int `gorm:"primaryKey,colum:question_id"`
	KeywordId  int `gorm:"primaryKey,colum:keyword_id"`
}
type SubjectiveQuestionKeywords struct {
	QuestionId int `gorm:"primaryKey,colum:question_id"`
	KeywordId  int `gorm:"primaryKey,colum:keyword_id"`
}

type Tests struct {
	Id         int `gorm:"primaryKey"`
	Name       string
	QuestionId int
	Grade      float64
	Author     string
}
type QuestionSummary struct {
	ID           int    `json:"id"`
	QuestionType string `json:"question_type"`
	Subject      string `json:"subject"`
	Content      string `json:"content"`
	Options      string `json:"options,omitempty"` // 选择题有值
	Difficulty   int    `json:"difficulty"`
	Author       string `json:"author"`
}

type Assignments struct {
	TestId     int    `gorm:"primaryKey"`
	QuestionId int    `gorm:"primaryKey"`
	StuName    string `gorm:"primaryKey"`
	Score      float64
	StuAnswer  string
	StuScore   float64
	AssignName string
}

func getDatabase() (*gorm.DB, error) {

	log.Println("Connecting to database")
	dsn := DatabaseUserName + ":" + DatabasePassword + "@tcp(" + DatabaseAddress + ")/" + DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Successfully connected to database")
	return db, nil

}
