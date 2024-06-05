package api

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DatabaseUserName = "root"
	DatabasePassword = "123456"

	DatabaseName = "SEProject"
)

func getDatabase() (*gorm.DB, error) {

	dsn := DatabaseUserName + ":" + DatabasePassword + "@tcp(localhost:3306)/" + DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}
