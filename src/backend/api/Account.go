package api

import (
	"gorm.io/gorm"
)

const (
	STUDENT = "STUDENT"
	TEACHER = "TEACHER"
	ADMIN   = "ADMIN"
)

type Users struct {
	Username string `gorm:"primaryKey"`
	Password string
	Type     string
}

func GetUserByUsername(db *gorm.DB, username string, user *Users) error {
	return db.Where("username = ?", username).First(user).Error
}

func AddUser(db *gorm.DB, user *Users) error {
	return db.Create(user).Error
}

func UpdateUser(db *gorm.DB, user *Users) error {
	return db.Save(user).Error
}

func DeleteUser(db *gorm.DB, username string) error {
	return db.Delete(&Users{}, "username = ?", username).Error
}
