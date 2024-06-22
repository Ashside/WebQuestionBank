package api

import "gorm.io/gorm"

func GetAssignsByAssignName(db *gorm.DB, assignName string) ([]Assignments, error) {
	var assign []Assignments
	if err := db.Table("assignments").Where("assign_name = ?", assignName).Find(&assign).Error; err != nil {
		return assign, err
	}
	return assign, nil
}

func GetStuAnswerByStuName(db *gorm.DB, stuName string) (Assignments, error) {
	var assign Assignments
	if err := db.Table("assignments").Where("stu_name = ?", stuName).Find(&assign).Error; err != nil {
		return assign, err
	}
	return assign, nil
}

func (a *Assignments) AddAssign(db *gorm.DB) error {
	err := db.Table("assignments").Create(a).Error
	if err != nil {
		return err
	}
	return nil
}
