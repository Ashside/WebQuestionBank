package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const AccessToken = "24.dcc9972f075b3221c78daffd648b02d8.2592000.1720265395.282335-79244792"

type keywordResponse struct {
	keyword string
	score   float64
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
			Score   float64 `json:"score"`
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
