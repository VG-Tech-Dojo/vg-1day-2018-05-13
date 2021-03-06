package bot

import (
	"regexp"
	"strings"

	"fmt"

	"net/url"

	"github.com/VG-Tech-Dojo/vg-1day-2018-05-13/mochiya98/env"
	"github.com/VG-Tech-Dojo/vg-1day-2018-05-13/mochiya98/model"
)

const (
	keywordAPIURLFormat = "https://jlp.yahooapis.jp/KeyphraseService/V1/extract?appid=%s&sentence=%s&output=json"
	talkAPIURLEndpoint  = "https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk"
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

	// GachaProcessor は"SSレア", "Sレア", "レア", "ノーマル"のいずれかをランダムで作るprocessorの構造体です
	GachaProcessor struct{}

	// TalkProcessor はメッセージ本文からキーワードを抽出するprocessorの構造体です
	TalkProcessor struct{}
)

// Process は"hello, world!"というbodyがセットされたメッセージのポインタを返します
func (p *HelloWorldProcessor) Process(msgIn *model.Message) (*model.Message, error) {
	return &model.Message{
		Body:     msgIn.Body + ", world!",
		UserName: "HelloBot",
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
		Body:     result,
		UserName: "OmikujiBot",
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
		Body:     "キーワード：" + strings.Join(keywords, ", "),
		UserName: "KeywordBot",
	}, nil
}

// Process は"SSレア", "Sレア", "レア", "ノーマル"のいずれかがbodyにセットされたメッセージへのポインタを返します
func (p *GachaProcessor) Process(msgIn *model.Message) (*model.Message, error) {
	fortunes := []string{
		"SSレア",
		"Sレア",
		"レア",
		"ノーマル",
	}
	result := fortunes[randIntn(len(fortunes))]
	return &model.Message{
		Body:     result,
		UserName: "GachaBot",
	}, nil
}

// Process はメッセージ本文からキーワードを抽出します
func (p *TalkProcessor) Process(msgIn *model.Message) (*model.Message, error) {
	r := regexp.MustCompile("\\Atalk (.*)\\z")
	matchedStrings := r.FindStringSubmatch(msgIn.Body)
	text := matchedStrings[1]

	type talkAPIResponseResult struct {
		Perplexity float64 `json:"perplexity"`
		Reply      string  `json:"reply"`
	}
	type talkAPIResponse struct {
		Status  int64                   `json:"status"`
		Message string                  `json:"message"`
		Results []talkAPIResponseResult `json:"results"`
	}
	var json talkAPIResponse

	requestBody := url.Values{}
	requestBody.Add("apikey", env.TalkAPIApiKey)
	requestBody.Add("query", text)

	post(talkAPIURLEndpoint, requestBody, &json)

	replies := []string{}
	for _, element := range json.Results {
		replies = append(replies, element.Reply)
	}

	return &model.Message{
		Body:     "返事: " + strings.Join(replies, ", "),
		UserName: "TalkBot",
	}, nil
}
