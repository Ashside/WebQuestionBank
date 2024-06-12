package api

import (
	"gorm.io/gorm"
	"log"
)

type ChoiceOptions struct {
	Option1 string `json:"option1"`
	Option2 string `json:"option2"`
	Option3 string `json:"option3"`
	Option4 string `json:"option4"`
}

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
func AddChoiceQuestion(db *gorm.DB, c *ChoiceQuestions) interface{} {
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
	if err != nil {
		log.Printf("Failed to query choiceQuestions: %v\n", err)
	} else {
		log.Println("Successfully queried choiceQuestions")
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
	return choiceQuestion, nil
}
