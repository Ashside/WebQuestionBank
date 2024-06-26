package api

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func findAvailableTestsId(db *gorm.DB) int {
	// 将id置为最大值加一
	var id int
	var maxId int64
	// 查询数据项条数
	if err := db.Table("tests").Count(&maxId).Error; err != nil {
		log.Println("Failed to count tests")
		return -1
	}
	if maxId == 0 {
		return 1
	}

	if err := db.Table("tests").Select("max(id)").Find(&maxId).Error; err != nil {
		log.Println("Failed to get max id")
		return -1
	}
	id = int(maxId + 1)
	return id
}
func GenerateMdByTestID(db *gorm.DB, id int) (string, error) {
	// 先查询所有题目
	var tests []Tests
	if err := db.Table("tests").Where("id = ?", id).Find(&tests).Error; err != nil {
		return "", err
	}
	var mdFile string
	log.Println(tests)
	for i, test := range tests {
		quesId := test.QuestionId
		// 查询题目
		ques, bExist := QueryQuestionFromId(db, quesId)
		if bExist {
			mdFile += "# 第" + strconv.Itoa(i+1) + "题\n"
			mdFile += strconv.Itoa(int(test.Grade)) + "分\n"
			mdFile += ques.Content + "\n"
			if ques.Options != "" {
				mdFile += "\n"
				mdFile += "选项：\n\n"

				var options map[string]string
				if err := json.Unmarshal([]byte(ques.Options), &options); err != nil {
					log.Println("Failed to unmarshal options")
					return "", err
				}
				for k, v := range options {
					mdFile += k + " : " + v + "\n"
					mdFile += "\n"
				}

			}
		} else {
			return "", nil
		}

	}
	return mdFile, nil
}
func GenerateMdByQuestions(db *gorm.DB, questions []QuestionSummary) (string, error) {
	// 生成md文件
	var mdFile string
	for i, ques := range questions {
		mdFile += "# 第" + strconv.Itoa(i+1) + "题\n"
		mdFile += ques.Content + "\n"
		if ques.Options != "" {
			mdFile += "\n"
			mdFile += "选项：\n\n"

			var options map[string]string
			if err := json.Unmarshal([]byte(ques.Options), &options); err != nil {
				log.Println("Failed to unmarshal options")
				return "", err
			}
			for k, v := range options {
				mdFile += k + " : " + v + "\n"
				mdFile += "\n"
			}

		}
	}
	return mdFile, nil

}

func AddTest(db *gorm.DB, t *Tests) error {
	// 查看name是否重复
	var tests []Tests
	if err := db.Table("tests").Where("name = ?", t.Name).Find(&tests).Error; err != nil {
		return err

	}
	// 如果有多个id不同的test的name相同，那么name+id
	for _, test := range tests {
		if test.Name == t.Name && test.Id != t.Id {
			t.Name += strconv.Itoa(t.Id)
		}

	}
	// 添加测试
	err := db.Table("tests").Create(t).Error
	if err != nil {
		return err
	}
	return nil

}

func GeneratePDFFile(file string, id int) (string, error) {
	// TODO 待完成
	//
	return "", nil
}

func QueryAllTests(db *gorm.DB, username string, userType string) ([]Tests, error) {
	// 查询所有测试
	var tests []Tests

	// 如果是管理员，查询所有测试
	if userType == ADMIN {
		if err := db.Table("tests").Find(&tests).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.Table("tests").Where("author = ?", username).Find(&tests).Error; err != nil {
			return nil, err
		}
	}
	return tests, nil
}
func QueryQuesIdByTestID(db *gorm.DB, id int) []int {
	// 返回该测试下所有题目的id

	var test []Tests
	var quesId []int
	if err := db.Table("tests").Where("id = ?", id).Find(&test).Error; err != nil {
		return nil
	}
	for _, t := range test {
		quesId = append(quesId, t.QuestionId)
	}
	return quesId
}

func QueryGradeByTestIdAndQuestionId(db *gorm.DB, testId int, questionId int) (float64, error) {
	// 查询该题目的分数
	var test Tests
	if err := db.Table("tests").Where("id = ? AND question_id = ?", testId, questionId).Find(&test).Error; err != nil {
		return 0, err
	}
	return test.Grade, nil
}

func QueryTestsByStuName(db *gorm.DB, username string) ([]Tests, error) {
	// 查询学生的所有测试
	var tests []Tests
	if err := db.Table("tests").Where("stu_name = ?", username).Find(&tests).Error; err != nil {
		return nil, err
	}
	return tests, nil
}
func QueryTestNameByTestId(db *gorm.DB, t int) string {
	var test Tests
	if err := db.Table("tests").Where("id = ?", t).First(&test).Error; err != nil {
		return ""
	}
	return test.Name
}

func isTestFinished(db *gorm.DB, t int, username string) bool {
	// 在assignments表中查询该学生是否完成了该测试
	var assign []Assignments
	if err := db.Table("assignments").Where("test_id = ? AND stu_name = ?", t, username).Find(&assign).Error; err != nil {
		return false
	}
	for _, a := range assign {
		if a.Finished == false {
			return false
		}

	}
	return true
}
