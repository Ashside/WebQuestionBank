package api

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	DatabaseUserName = "root"
	DatabasePassword = "Aa=12345678"
	DatabaseName     = "SEProject"

	DatabaseAddress = "121.43.124.218:3306"
)

type Users struct {
	Username string `gorm:"primaryKey"`
	Password string
	Type     string
}
type Keywords struct {
	Id      int `gorm:"primaryKey"`
	Keyword string
	Score   float64
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
