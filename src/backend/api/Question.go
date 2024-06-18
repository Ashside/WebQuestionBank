package api

import (
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
func QueryQuestionFromId(db *gorm.DB, id int) (ChoiceQuestions, error) {
	var choiceQuestion ChoiceQuestions
	var subjectiveQuestion SubjectiveQuestions
	if id == 0 {
		log.Println("Empty id")
		return choiceQuestion, nil
	}
	// 查询ChoiceQuestions表中的题目
	err := db.Table("choicequestions").Where("id = ?", id).Find(&choiceQuestion).Error
	if err != nil || choiceQuestion.Id == 0 {
		log.Printf("Failed to query choiceQuestions: %v\n", err)
	} else {
		log.Println("Successfully queried choiceQuestions")
		log.Println(choiceQuestion.Id)
	}
	// 查询SubjectiveQuestions表中的题目
	err = db.Table("subjectivequestions").Where("id = ?", id).Find(&subjectiveQuestion).Error
	if err != nil {
		log.Printf("Failed to query subjectiveQuestions: %v\n", err)
	} else {
		log.Println("Successfully queried subjectiveQuestions")
	}

	if choiceQuestion.Id == 0 {
		choiceQuestion = subjectiveQuestion.ToChoiceQuestions()
	}
	if err != nil {
		log.Printf("Failed to query question: %v\n", err)

		return ChoiceQuestions{}, err
	} else {
		log.Println("Successfully queried question")

	}
	return choiceQuestion, nil
}

func DeleteChoiceQuestion(db *gorm.DB, id int) error {
	// 需要确保id存在
	// 删除ChoiceQuestions表中的题目

	// 先删除与该题目相关的关键词

	// 查询choice_question_keywords表中的关键词
	// 同时根据keyword_id删除keywords表中的关键词
	var keywords []Keywords
	err := db.Table("choice_question_keywords").Where("question_id = ?", id).Find(&keywords).Error
	if err != nil {
		log.Printf("Failed to query keywords: %v\n", err)
	} else {
		log.Println("Successfully queried keywords")
	}

	// 删除choice_question_keywords表中的关键词
	err = db.Table("choice_question_keywords").Where("question_id = ?", id).Delete(&choiceQuestionKeywords{}).Error
	if err != nil {
		log.Printf("Failed to delete keyword: %v\n", err)
	} else {
		log.Println("Successfully deleted keyword")
	}

	for _, keyword := range keywords {
		err = db.Table("keywords").Where("id = ?", keyword.Id).Delete(&Keywords{}).Error
		if err != nil {
			log.Printf("Failed to delete keyword: %v\n", err)
		} else {
			log.Println("Successfully deleted keyword")
		}

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

	// 先删除与该题目相关的关键词

	// 查询subjective_question_keywords表中的关键词
	// 同时根据keyword_id删除keywords表中的关键词
	var keywords []Keywords
	err := db.Table("subjective_question_keywords").Where("question_id = ?", id).Find(&keywords).Error
	if err != nil {
		log.Printf("Failed to query keywords: %v\n", err)
	} else {
		log.Println("Successfully queried keywords")
	}

	// 删除subjective_question_keywords表中的关键词
	err = db.Table("subjective_question_keywords").Where("question_id = ?", id).Delete(&subjectiveQuestionKeywords{}).Error
	if err != nil {
		log.Printf("Failed to delete keyword: %v\n", err)
	} else {
		log.Println("Successfully deleted keyword")
	}

	for _, keyword := range keywords {
		err = db.Table("keywords").Where("id = ?", keyword.Id).Delete(&Keywords{}).Error
		if err != nil {
			log.Printf("Failed to delete keyword: %v\n", err)
		} else {
			log.Println("Successfully deleted keyword")
		}
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

	fromId, err := QueryQuestionFromId(db, int(id))
	if err != nil {
		return false
	}
	// 如果id为0，说明不存在，虽然id为0的题目理论上也不应该存在
	if fromId.Id == 0 {
		return false
	}
	return true

}

func findAvailableID(db *gorm.DB) int {
	// 从1开始查找可用的id
	for i := 1; ; i++ {
		if !isQuestionExistFromID(db, int64(i)) {
			return i
		}
	}
}
