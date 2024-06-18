package api

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const AccessToken = "24.dcc9972f075b3221c78daffd648b02d8.2592000.1720265395.282335-79244792"

type keywordResponse struct {
	Keyword string
	Score   float64
}

func getKeyword(text string) ([]keywordResponse, error) {

	url := "https://aip.baidubce.com/rpc/2.0/nlp/v1/txt_keywords_extraction?access_token=" + AccessToken
	payload := strings.NewReader(`{"text":["` + text + `"]}`)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 分析返回的json数据
	// 创建一个用于接收JSON数据的结构体
	var result struct {
		Items []struct {
			Keyword string  `json:"word"`
			Score   float64 `json:"Score"`
		} `json:"results"`
	}

	// 解析JSON数据
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 打印关键词
	for _, item := range result.Items {
		log.Println(item.Keyword, item.Score)
	}
	// 返回关键词
	var response []keywordResponse
	for _, item := range result.Items {
		response = append(response, keywordResponse{item.Keyword, item.Score})
	}
	return response, nil

}
func findAvailableKeywordId(db *gorm.DB) int {
	// 查询数据库中的关键词个数
	// 如果关键词个数为0，则将id设置为1
	// 否则将id设置为关键词个数最近的一个可用id
	// 返回id
	var id int64
	// 查询数据库中的关键词个数
	if err := db.Table("keywords").Count(&id).Error; err != nil {
		log.Println("Failed to count keywords")
		return -1
	}
	// 如果关键词个数为0，则将id设置为1
	if id == 0 {
		id = 1
	} else {
		// 否则从id开始递增
		for cnt := id; cnt > 0; cnt++ {
			if err := db.Table("keywords").Where("id = ?", cnt).Find(&Keywords{}).Error; err != nil {
				// 如果找不到id为cnt的关键词，则将id设置为cnt
				id = cnt
				break
			}
		}
	}
	log.Println("Find Keyword id:", id)
	return int(id)
}

func genKeywordId(db *gorm.DB, keyword string) (int, bool) {
	// 先查找是否有相同的关键词
	var existingKeyword Keywords
	if err := db.Table("keywords").Where("keyword = ?", keyword).First(&existingKeyword).Error; err == nil {
		// 如果有则返回该关键词的id
		log.Println("Existing Keyword id:", existingKeyword.Id)
		return existingKeyword.Id, true
	}

	// 否则调用findAvailableKeywordId函数查找可用的id
	log.Println("Generating Keyword id")
	id := findAvailableKeywordId(db)
	return id, false
}

func AddKeywords(db *gorm.DB, keywords []keywordResponse, quesId int, isChoice bool) {

	for _, keyword := range keywords {
		// 生成关键词id
		id, isExist := genKeywordId(db, keyword.Keyword)
		if id == -1 {
			log.Println("Failed to generate Keyword id")
			return
		}
		// 如果关键词已经存在，则不需要添加关键词
		if isExist {
			log.Println("Keyword already exists")

		} else {
			// 添加关键词
			log.Println("Adding Keyword:", keyword.Keyword)
			err := db.Table("keywords").Create(&Keywords{Id: id, Keyword: keyword.Keyword}).Error
			if err != nil {
				log.Println("Failed to add Keyword:", keyword.Keyword)
				return
			}
			log.Println("Successfully added Keyword:", keyword.Keyword)
		}

		// 添加关键词和题目的关系
		log.Println("Adding Keyword-question relationship")
		if isChoice {
			err := db.Table("choice_question_keywords").Create(&ChoiceQuestionKeywords{QuestionId: quesId, KeywordId: id}).Error
			if err != nil {
				log.Println("Failed to add Keyword-question relationship")
				log.Println(err.Error())
				return

			}
		} else {
			err := db.Table("subjective_question_keywords").Create(&SubjectiveQuestionKeywords{QuestionId: quesId, KeywordId: id}).Error
			if err != nil {
				log.Println("Failed to add Keyword-question relationship")
				log.Println(err.Error())
				return
			}
		}

	}
	log.Println("Successfully added keywords")

}

func GetKeywordsByQuestionId(db *gorm.DB, id int, bChoiceQues bool) ([]Keywords, error) {
	var keywords []Keywords
	if bChoiceQues {
		err := db.Table("keywords").Joins("JOIN choice_question_keywords ON choice_question_keywords.keyword_id = keywords.id").Where("choice_question_keywords.question_id = ?", id).Find(&keywords).Error
		if err != nil {
			log.Println("Failed to get keywords")
			return nil, err
		}
	} else {
		err := db.Table("keywords").Joins("JOIN subjective_question_keywords ON subjective_question_keywords.keyword_id = keywords.id").Where("subjective_question_keywords.question_id = ?", id).Find(&keywords).Error
		if err != nil {
			log.Println("Failed to get keywords")
			return nil, err
		}
	}
	log.Println("Successfully get keywords")
	return keywords, nil

}
