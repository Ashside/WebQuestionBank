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
type choiceQuestionKeywords struct {
	QuestionId int `gorm:"primaryKey,colum:question_id"`
	KeywordId  int `gorm:"primaryKey,colum:keyword_id"`
}
type subjectiveQuestionKeywords struct {
	QuestionId int `gorm:"primaryKey,colum:question_id"`
	KeywordId  int `gorm:"primaryKey,colum:keyword_id"`
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
func genKeywordId(db *gorm.DB) int {
	// 查询数据库中的关键词个数
	// 如果关键词个数为0，则将id设置为1
	// 否则将id设置为关键词个数+1
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
		// 否则将id设置为关键词个数+1
		id++
	}
	log.Println("Generated Keyword id:", id)
	return int(id)

}
func AddKeywords(db *gorm.DB, keywords []keywordResponse, quesId int, isChoice bool) {

	for _, keyword := range keywords {
		// TODO 特异性处理
		// 生成关键词id
		id := genKeywordId(db)
		if id == -1 {
			log.Println("Failed to generate Keyword id")
			return
		}

		// 添加关键词
		log.Println("Adding Keyword:", keyword.Keyword)
		err := db.Table("keywords").Create(&Keywords{Id: id, Keyword: keyword.Keyword, Score: keyword.Score}).Error
		if err != nil {
			log.Println("Failed to add Keyword:", keyword.Keyword)
			return
		}
		log.Println("Successfully added Keyword:", keyword.Keyword)

		// 添加关键词和题目的关系
		log.Println("Adding Keyword-question relationship")
		if isChoice {
			err = db.Table("choice_question_keywords").Create(&choiceQuestionKeywords{QuestionId: quesId, KeywordId: id}).Error
			if err != nil {
				log.Println("Failed to add Keyword-question relationship")
				println(err.Error())
				return

			}
		} else {
			err = db.Table("subjective_question_keywords").Create(&subjectiveQuestionKeywords{QuestionId: quesId, KeywordId: id}).Error
			if err != nil {
				log.Println("Failed to add Keyword-question relationship")
				println(err.Error())
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
