package api

import (
	"gorm.io/gorm"
	"log"
)

// AddSubjectQuestion 向数据库中添加主观题，打印log信息
func AddSubjectQuestion(db *gorm.DB, question *SubjectiveQuestions) error {
	log.Println("Adding question", question)

	// 向SubjectiveQuestions表中添加题目
	err := db.Table("SubjectiveQuestions").Create(question).Error
	if err != nil {
		log.Printf("Failed to add question: %v\n", err)
	} else {
		log.Println("Successfully added question")
	}

	return err
}
func AddChoiceQuestion(db *gorm.DB, c *ChoiceQuestions) interface{} {
	log.Println("Adding question", c)

	// 向ChoiceQuestions表中添加题目
	err := db.Table("ChoiceQuestions").Create(c).Error
	if err != nil {
		log.Printf("Failed to add question: %v\n", err)
	} else {
		log.Println("Successfully added question")
	}

	return err
}
func QueryQuestion(db *gorm.DB, username string, subject string, difficulty int) []ChoiceQuestions {
	var choiceQuestions []ChoiceQuestions
	var subjectiveQuestions []ChoiceQuestions
	// 查询ChoiceQuestions表中的题目
	err := db.Table("ChoiceQuestions").Where("author = ? AND subject = ? AND difficulty = ?", username, subject, difficulty).Find(&choiceQuestions).Error
	if err != nil {
		log.Printf("Failed to query choiceQuestions: %v\n", err)
	} else {
		log.Println("Successfully queried choiceQuestions")
	}
	// 查询SubjectiveQuestions表中的题目
	err = db.Table("SubjectiveQuestions").Where("author = ? AND subject = ? AND difficulty = ?", username, subject, difficulty).Find(&subjectiveQuestions).Error
	if err != nil {
		log.Printf("Failed to query subjectiveQuestions: %v\n", err)
	} else {
		log.Println("Successfully queried subjectiveQuestions")
	}
	// 将subjectiveQuestions转换为ChoiceQuestions
	for _, q := range subjectiveQuestions {
		choiceQuestions = append(choiceQuestions, ChoiceQuestions{
			Id:         q.Id,
			Subject:    q.Subject,
			Content:    q.Content,
			Options:    "",
			Answer:     q.Answer,
			Difficulty: q.Difficulty,
		})
	}

	return choiceQuestions
}
