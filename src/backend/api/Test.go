package api

import (
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
func GenerateMD(db *gorm.DB, id int) (string, error) {
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
			mdFile += "第" + strconv.Itoa(i+1) + "题\n"
			mdFile += strconv.Itoa(int(test.Grade)) + "分\n"
			mdFile += ques.Content + "\n"
			if ques.Options != "" {
				mdFile += "选项：" + ques.Options + "\n"
			}
		} else {
			return "", nil
		}

	}
	return mdFile, nil
}

func AddTest(db *gorm.DB, t *Tests) error {
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
