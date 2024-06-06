package api

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	DatabaseUserName = "root"
	DatabasePassword = "123456"
	DatabaseName     = "SEProject"
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

func getDatabase() (*gorm.DB, error) {

	log.Println("Connecting to database")
	dsn := DatabaseUserName + ":" + DatabasePassword + "@tcp(localhost:3306)/" + DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}
