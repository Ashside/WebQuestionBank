package api

import (
	"gorm.io/gorm"
	"log"
)

const (
	STUDENT = "student"
	TEACHER = "teacher"
	ADMIN   = "admin"
)

func GetUserByUsername(db *gorm.DB, username string, user *Users) error {

	log.Printf("Get user: %s\n", username)
	err := db.Where("username = ?", username).First(user).Error
	if err != nil {
		log.Printf("Failed to get user: %v\n", err)
	} else {
		log.Println("Successfully get user")
	}
	return err
}

func AddUser(db *gorm.DB, user *Users) error {
	log.Printf("Adding user: %+v\n", *user)
	err := db.Create(user).Error
	if err != nil {
		log.Printf("Failed to add user: %v\n", err)
	} else {
		log.Println("Successfully added user")
	}
	return err
}

func UpdateUser(db *gorm.DB, user *Users) error {
	return db.Save(user).Error
}

func DeleteUser(db *gorm.DB, username string) error {
	return db.Delete(&Users{}, "username = ?", username).Error
}

func (user *Users) IsStudent() bool {
	return user.Type == STUDENT
}

func (user *Users) IsTeacher() bool {
	return user.Type == TEACHER
}

func (user *Users) IsAdmin() bool {
	return user.Type == ADMIN
}
