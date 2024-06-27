package api

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

// AddSubjectQuestion 向数据库中添加主观题，打印log信息
func AddSubjectQuestion(db *gorm.DB, question *SubjectiveQuestions) error {
	log.Println("Adding question", question)

	// 向SubjectiveQuestions表中添加题目
	err := db.Table("subjectivequestions").Create(question).Error
	if err != nil {
		log.Printf("Failed to add question: %v\n", err)
	} else {
		log.Println("Successfully added question")
	}

	return err
}
func AddChoiceQuestion(db *gorm.DB, c *ChoiceQuestions) error {
	log.Println("Adding question", c)

	// 向ChoiceQuestions表中添加题目
	err := db.Table("choicequestions").Create(c).Error
	if err != nil {
		log.Printf("Failed to add question: %v\n", err)
	} else {
		log.Println("Successfully added question")
	}

	return err
}
func QueryQuestionFromCertainInf(db *gorm.DB, username string, subject string, difficulty int) []ChoiceQuestions {
	var choiceQuestions []ChoiceQuestions
	var subjectiveQuestions []SubjectiveQuestions

	/*
		// 查询ChoiceQuestions表中的题目
		err := db.Table("choicequestions").Where("author = ? AND subject = ? AND difficulty = ?", username, subject, difficulty).Find(&choiceQuestions).Error
		if err != nil {
			log.Printf("Failed to query choiceQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried choiceQuestions")
		}
		// 查询SubjectiveQuestions表中的题目
		err = db.Table("subjectivequestions").Where("author = ? AND subject = ? AND difficulty = ?", username, subject, difficulty).Find(&subjectiveQuestions).Error
		if err != nil {
			log.Printf("Failed to query subjectiveQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried subjectiveQuestions")
		}*/

	// 根据传入的参数查询ChoiceQuestions表中的题目
	// 先检查参数是否给出
	if username == "" && subject == "" && difficulty == 0 {
		// 从表中查询所有题目
		err := db.Table("choicequestions").Find(&choiceQuestions).Error

		if err != nil {
			log.Printf("Failed to query choiceQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried choiceQuestions")
		}
		err = db.Table("subjectivequestions").Find(&subjectiveQuestions).Error
		if err != nil {
			log.Printf("Failed to query subjectiveQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried subjectiveQuestions")
		}

	}

	// 根据传入的参数查询ChoiceQuestions表中的题目
	if username != "" && subject != "" && difficulty != 0 {
		err := db.Table("choicequestions").Where("author = ? AND subject = ? AND difficulty = ?", username, subject, difficulty).Find(&choiceQuestions).Error
		if err != nil {
			log.Printf("Failed to query choiceQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried choiceQuestions")
		}
		err = db.Table("subjectivequestions").Where("author = ? AND subject = ? AND difficulty = ?", username, subject, difficulty).Find(&subjectiveQuestions).Error
		if err != nil {
			log.Printf("Failed to query subjectiveQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried subjectiveQuestions")
		}
	}

	if username != "" && subject == "" && difficulty == 0 {
		err := db.Table("choicequestions").Where("author = ?", username).Find(&choiceQuestions).Error
		if err != nil {
			log.Printf("Failed to query choiceQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried choiceQuestions")
		}
		err = db.Table("subjectivequestions").Where("author = ?", username).Find(&subjectiveQuestions).Error
		if err != nil {
			log.Printf("Failed to query subjectiveQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried subjectiveQuestions")
		}
	}

	if username == "" && subject != "" && difficulty == 0 {
		err := db.Table("choicequestions").Where("subject = ?", subject).Find(&choiceQuestions).Error
		if err != nil {
			log.Printf("Failed to query choiceQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried choiceQuestions")
		}
		err = db.Table("subjectivequestions").Where("subject = ?", subject).Find(&subjectiveQuestions).Error
		if err != nil {
			log.Printf("Failed to query subjectiveQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried subjectiveQuestions")
		}

	}

	if username == "" && subject == "" && difficulty != 0 {
		err := db.Table("choicequestions").Where("difficulty = ?", difficulty).Find(&choiceQuestions).Error
		if err != nil {
			log.Printf("Failed to query choiceQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried choiceQuestions")
		}
		err = db.Table("subjectivequestions").Where("difficulty = ?", difficulty).Find(&subjectiveQuestions).Error
		if err != nil {
			log.Printf("Failed to query subjectiveQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried subjectiveQuestions")
		}

	}

	if username != "" && subject != "" && difficulty == 0 {
		err := db.Table("choicequestions").Where("author = ? AND subject = ?", username, subject).Find(&choiceQuestions).Error
		if err != nil {
			log.Printf("Failed to query choiceQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried choiceQuestions")
		}
		err = db.Table("subjectivequestions").Where("author = ? AND subject = ?", username, subject).Find(&subjectiveQuestions).Error
		if err != nil {
			log.Printf("Failed to query subjectiveQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried subjectiveQuestions")
		}
	}

	if username != "" && subject == "" && difficulty != 0 {
		err := db.Table("choicequestions").Where("author = ? AND difficulty = ?", username, difficulty).Find(&choiceQuestions).Error
		if err != nil {
			log.Printf("Failed to query choiceQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried choiceQuestions")
		}
		err = db.Table("subjectivequestions").Where("author = ? AND difficulty = ?", username, difficulty).Find(&subjectiveQuestions).Error
		if err != nil {
			log.Printf("Failed to query subjectiveQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried subjectiveQuestions")
		}
	}

	if username == "" && subject != "" && difficulty != 0 {
		err := db.Table("choicequestions").Where("subject = ? AND difficulty = ?", subject, difficulty).Find(&choiceQuestions).Error
		if err != nil {
			log.Printf("Failed to query choiceQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried choiceQuestions")
		}
		err = db.Table("subjectivequestions").Where("subject = ? AND difficulty = ?", subject, difficulty).Find(&subjectiveQuestions).Error
		if err != nil {
			log.Printf("Failed to query subjectiveQuestions: %v\n", err)
		} else {
			log.Println("Successfully queried subjectiveQuestions")
		}
	}

	// 将subjectiveQuestions转换为ChoiceQuestions
	for _, q := range subjectiveQuestions {
		choiceQuestions = append(choiceQuestions, q.ToChoiceQuestions())
	}

	return choiceQuestions
}
func (s *SubjectiveQuestions) ToChoiceQuestions() ChoiceQuestions {
	return ChoiceQuestions{
		Id:         s.Id,
		Subject:    s.Subject,
		Content:    s.Content,
		Options:    "",
		Answer:     s.Answer,
		Difficulty: s.Difficulty,
	}
}
func QueryQuestionFromId(db *gorm.DB, id int) (ChoiceQuestions, bool) {
	var choiceQuestion ChoiceQuestions
	var subjectiveQuestion SubjectiveQuestions
	if id == 0 {
		log.Println("Empty id")
		return choiceQuestion, false
	}
	// 查询ChoiceQuestions表中的题目
	err := db.Table("choicequestions").Where("id = ?", id).Find(&choiceQuestion).Error
	if err != nil || choiceQuestion.Id == 0 {
		log.Printf("Failed to query choiceQuestions: %v\n", err)

	} else {
		log.Println("Successfully queried choiceQuestions")
		log.Println(choiceQuestion.Id)
		return choiceQuestion, true
	}
	// 查询SubjectiveQuestions表中的题目
	err = db.Table("subjectivequestions").Where("id = ?", id).Find(&subjectiveQuestion).Error
	if err != nil || subjectiveQuestion.Id == 0 {
		log.Printf("Failed to query subjectiveQuestions: %v\n", err)
		// 这里查询失败一定意味着id不存在
		return ChoiceQuestions{}, false
	} else {
		log.Println("Successfully queried subjectiveQuestions")
		choiceQuestion = subjectiveQuestion.ToChoiceQuestions()
		return choiceQuestion, true
	}

}

func DeleteChoiceQuestion(db *gorm.DB, id int) error {
	// 需要确保id存在
	// 删除ChoiceQuestions表中的题目

	// 删除choice_question_keywords表中的关键词
	err := db.Table("choice_question_keywords").Where("question_id = ?", id).Delete(&ChoiceQuestionKeywords{}).Error
	if err != nil {
		log.Printf("Failed to delete keyword: %v\n", err)
	} else {
		log.Println("Successfully deleted keyword")
	}

	err = db.Table("choicequestions").Where("id = ?", id).Delete(&ChoiceQuestions{}).Error
	if err != nil {
		log.Printf("Failed to delete question: %v\n", err)
	} else {
		log.Println("Successfully deleted question")
	}

	return err

}

func DeleteSubjectQuestion(db *gorm.DB, id int) error {
	// 需要确保id存在
	// 删除SubjectiveQuestions表中的题目

	// 删除subjective_question_keywords表中的关键词
	err := db.Table("subjective_question_keywords").Where("question_id = ?", id).Delete(&SubjectiveQuestionKeywords{}).Error
	if err != nil {
		log.Printf("Failed to delete keyword: %v\n", err)
	} else {
		log.Println("Successfully deleted relationship")
	}

	err = db.Table("subjectivequestions").Where("id = ?", id).Delete(&SubjectiveQuestions{}).Error
	if err != nil {
		log.Printf("Failed to delete question: %v\n", err)
	} else {
		log.Println("Successfully deleted question")
	}

	return err

}

func isQuestionExistFromID(db *gorm.DB, id int64) bool {

	fromId, bExist := QueryQuestionFromId(db, int(id))
	if bExist != true {
		return false
	}
	// 如果id为0，说明不存在，虽然id为0的题目理论上也不应该存在
	if fromId.Id == 0 {
		return false
	}
	return true

}

func findAvailableQuesID(db *gorm.DB) int {
	// 从1开始查找可用的id
	for i := 1; ; i++ {
		if !isQuestionExistFromID(db, int64(i)) {
			return i
		}
	}
}

func getQuestionTypeByID(db *gorm.DB, id int) (string, error) {
	// 尝试查询选择题
	var choice QuestionSummary
	if err := db.Table("choicequestions").Where("id = ?", id).Take(&choice).Error; err == nil {
		return "choice", nil
	}

	// 如果不是选择题，则尝试查询主观题
	var subjective QuestionSummary
	if err := db.Table("subjectivequestions").Where("id = ?", id).Take(&subjective).Error; err == nil {
		return "subjective", nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	// 未知类型
	return "unknown", nil
}
func getQuestionsByTypeID(db *gorm.DB, questionType string, ids []int) ([]QuestionSummary, error) {
	var questions []QuestionSummary

	var tableName string
	var sql string
	switch questionType {
	case "Choice":
		tableName = "choicequestions"
		sql = "id, ? AS question_type, subject, content, options, difficulty, author"
	case "Subjective":
		tableName = "subjectivequestions"
		sql = "id, ? AS question_type, subject, content, difficulty, author"
	default:
		return nil, fmt.Errorf("unsupported question type: %s", questionType)
	}

	// 调整Select子句，使question_type字段动态反映表名
	if err := db.Table(tableName).
		Select(sql, tableName).
		Where("id IN (?)", ids).
		Scan(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func isChoiceQuestion(db *gorm.DB, id int) bool {
	var choice ChoiceQuestions
	if err := db.Table("choicequestions").Where("id = ?", id).Take(&choice).Error; err != nil {
		return false
	}
	return true
}

func isSubjectiveQuestion(db *gorm.DB, id int) bool {
	var subjective SubjectiveQuestions
	if err := db.Table("subjectivequestions").Where("id = ?", id).Take(&subjective).Error; err != nil {
		return false
	}
	return true
}
func isQuestionInTest(db *gorm.DB, id int) (bool, error) {
	var test Tests
	if err := db.Table("tests").Where("question_id = ?", id).First(&test).Error; err == nil {
		return true, nil
	}
	return false, nil
}
