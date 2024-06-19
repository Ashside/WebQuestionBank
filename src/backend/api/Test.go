package api

import (
	"gorm.io/gorm"
	"log"
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
func GeneratePDF(db *gorm.DB, id int) (string, error) {
	// 先查询所有题目
	var tests []Tests
	if err := db.Table("tests").Where("id = ?", id).Find(&tests).Error; err != nil {
		return "", err
	}
	log.Println(tests)
	return "", nil
}

func AddTest(db *gorm.DB, t *Tests) error {
	// 添加测试
	err := db.Table("tests").Create(t).Error
	if err != nil {
		return err
	}
	return nil

}
