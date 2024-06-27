package api

import (
	"gorm.io/gorm"
)

func GetAssignsByAssignName(db *gorm.DB, assignName string) ([]Assignments, error) {
	var assign []Assignments
	if err := db.Table("assignments").Where("assign_name = ?", assignName).Find(&assign).Error; err != nil {
		return assign, err
	}
	return assign, nil
}

func GetAssignsByStuName(db *gorm.DB, stuName string) ([]Assignments, error) {
	var assign []Assignments
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

func (a *Assignments) UpdateScore(db *gorm.DB) error {
	err := db.Table("assignments").Where("test_id = ? AND question_id = ? AND stu_name = ?", a.TestId, a.QuestionId, a.StuName).Update("stu_score", a.StuScore).Error
	if err != nil {
		return err
	}
	return nil
}

func QueryAssignsByTestAndStu(db *gorm.DB, testId int, stuName string) ([]Assignments, error) {
	var assign []Assignments
	if err := db.Table("assignments").Where("test_id = ? AND stu_name = ?", testId, stuName).Find(&assign).Error; err != nil {
		return assign, err
	}
	return assign, nil
}

func (a *Assignments) UpdateAnswer(db *gorm.DB) error {
	err := db.Table("assignments").Where("test_id = ? AND question_id = ? AND stu_name = ?", a.TestId, a.QuestionId, a.StuName).Update("stu_answer", a.StuAnswer).Error
	if err != nil {
		return err
	}
	return nil
}
func (a *Assignments) UpdateFinished(db *gorm.DB) error {
	err := db.Table("assignments").Where("test_id = ? AND question_id = ? AND stu_name = ?", a.TestId, a.QuestionId, a.StuName).Update("finished", a.Finished).Error
	if err != nil {
		return err
	}
	return nil
}
func CheckScore(db *gorm.DB, assign Assignments) float64 {
	bChoice := isChoiceQuestion(db, assign.QuestionId)
	if bChoice {
		// choice
		ques, _ := QueryQuestionFromId(db, assign.QuestionId)

		if ques.Answer == assign.StuAnswer {
			score, _ := QueryGradeByTestIdAndQuestionId(db, assign.TestId, assign.QuestionId)
			return score
		} else {
			return 0
		}
	} else {
		// subjective
		return -1
	}

}

func DeleteAssignByTestID(db *gorm.DB, testID int) error {
	err := db.Table("assignments").Where("test_id = ?", testID).Delete(&Assignments{}).Error
	if err != nil {
		return err
	}
	return nil
}
