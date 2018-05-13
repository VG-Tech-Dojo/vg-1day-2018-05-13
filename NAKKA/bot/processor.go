package bot

import (
	"regexp"
	"strings"

	"fmt"
	"log"

	"github.com/VG-Tech-Dojo/vg-1day-2018-05-13/NAKKA/env"
	"github.com/VG-Tech-Dojo/vg-1day-2018-05-13/NAKKA/model"
	"net/url"
)

const (
	keywordAPIURLFormat = "https://jlp.yahooapis.jp/KeyphraseService/V1/extract?appid=%s&sentence=%s&output=json"
	talkAPIURL          = "https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk"
)

type (
	// Processor はmessageを受け取り、投稿用messageを作るインターフェースです
	Processor interface {
		Process(message *model.Message) (*model.Message, error)
	}

	// HelloWorldProcessor は"hello, world!"メッセージを作るprocessorの構造体です
	HelloWorldProcessor struct{}

	// OmikujiProcessor は"大吉", "吉", "中吉", "小吉", "末吉", "凶"のいずれかをランダムで作るprocessorの構造体です
	OmikujiProcessor struct{}

	// KeywordProcessor はメッセージ本文からキーワードを抽出するprocessorの構造体です
	KeywordProcessor struct{}

	// GachaProcesseor はレア度のいずれかをランダムで作るprocessorの構造体です
	GachaProcessor struct{}

	TalkProcessor struct{}

	YoichiProcessor struct{}
)

// Process は"hello, world!"というbodyがセットされたメッセージのポインタを返します
func (p *HelloWorldProcessor) Process(msgIn *model.Message) (*model.Message, error) {
	return &model.Message{
		Body: msgIn.Body + ", world!",
	}, nil
}

// Process は"大吉", "吉", "中吉", "小吉", "末吉", "凶"のいずれかがbodyにセットされたメッセージへのポインタを返します
func (p *OmikujiProcessor) Process(msgIn *model.Message) (*model.Message, error) {
	fortunes := []string{
		"大吉",
		"吉",
		"中吉",
		"小吉",
		"末吉",
		"凶",
	}
	result := fortunes[randIntn(len(fortunes))]
	return &model.Message{
		Body: result,
	}, nil
}

// Process はメッセージ本文からキーワードを抽出します
func (p *KeywordProcessor) Process(msgIn *model.Message) (*model.Message, error) {
	r := regexp.MustCompile("\\Akeyword (.*)\\z")
	matchedStrings := r.FindStringSubmatch(msgIn.Body)
	text := matchedStrings[1]

	url := fmt.Sprintf(keywordAPIURLFormat, env.KeywordAPIAppID, url.QueryEscape(text))

	type keywordAPIResponse map[string]interface{}
	var json keywordAPIResponse
	get(url, &json)

	keywords := []string{}
	for k, v := range json {
		if k == "Error" {
			return nil, fmt.Errorf("%#v", v)
		}
		keywords = append(keywords, k)
	}

	return &model.Message{
		Body: "キーワード：" + strings.Join(keywords, ", "),
	}, nil
}

func (p *GachaProcessor) Process(msgIn *model.Message) (*model.Message, error) {
	fortunes := []string{
		"SSレア",
		"Sレア",
		"レア",
		"ノーマル",
	}
	result := fortunes[randIntn(len(fortunes))]
	return &model.Message{
		Body: result,
	}, nil
}

func (p *TalkProcessor) Process(msgIn *model.Message) (*model.Message, error) {
	r := regexp.MustCompile("\\Atalk (.*)\\z")
	matchedStrings := r.FindStringSubmatch(msgIn.Body)
	text := matchedStrings[1]

	res := &struct {
		Status  int64  `json:"status"`
		Message string `json:"message"`
		Results []struct {
			Perplexity float64 `json:"perplexity"`
			Reply      string  `json:"reply"`
		} `json:"results"`
	}{}

	params := url.Values{}
	params.Set("apikey", env.TalkAPIKey)
	params.Add("query", text)

	post(talkAPIURL, params, res)

	if res.Status != 0 {
		return nil, fmt.Errorf("%#v", res)
	}

	return &model.Message{
		Body:     "TalkAPI：" + res.Results[0].Reply,
		UserName: "talkbot",
	}, nil
}

func (p *YoichiProcessor) Process(msgIn *model.Message) (*model.Message, error) {
	r := regexp.MustCompile("\\Ayoichi (.*)\\z")
	matchedStrings := r.FindStringSubmatch(msgIn.Body)
	fmt.Printf("Process: %v\n", matchedStrings)
	text := matchedStrings[1]

	url := fmt.Sprintf(keywordAPIURLFormat, env.KeywordAPIAppID, url.QueryEscape(text))

	type keywordAPIResponse map[string]interface{}
	var json keywordAPIResponse
	get(url, &json)
	log.Printf("get: %v\n", json)

	keywords := []string{}
	for k, v := range json {
		if k == "Error" {
			return nil, fmt.Errorf("%#v", v)
		}
		keywords = append(keywords, k)
	}

	if len(keywords) == 0 {
		return &model.Message{
			Body:     "[No keyword!!]",
			UserName: "Yoichi Ochiai",
		}, nil
	}

	res, err := twitterGet(keywords[0])
	if err != nil {
		return nil, fmt.Errorf("%#v", err)
	}
	myBody := []string{res}

	return &model.Message{
		Body:     strings.Join(myBody, ", "),
		UserName: "Yoichi Ochiai",
	}, nil
}
