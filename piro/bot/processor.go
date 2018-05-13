package bot

import (
	"regexp"
	"strings"

	"fmt"

	"github.com/VG-Tech-Dojo/vg-1day-2018-05-13/piro/env"
	"github.com/VG-Tech-Dojo/vg-1day-2018-05-13/piro/model"

	"net/url"
	// "net/http"
  // "encoding/json"
)

const (
	keywordAPIURLFormat = "https://jlp.yahooapis.jp/KeyphraseService/V1/extract?appid=%s&sentence=%s&output=json"
	talkAPIURLFormat = "https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk"
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

	// GachaProcessor は"SSR", "SR", "R", "N"のいずれかをランダムで作るprocessorの構造体です
	GachaProcessor struct{}

	// ChatProcessor はメッセージ本文に対してAPIで返答を作るprocessorの構造体です
	TalkProcessor struct{}
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
		Username: "Omikuji Bot",
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
		Username: "Keyword Bot",
		Body: "キーワード：" + strings.Join(keywords, ", "),
	}, nil
}

// Process は"SSR", "SR", "R", "N"のいずれかがbodyにセットされたメッセージへのポインタを返します
func (p *GachaProcessor) Process(msgIn *model.Message) (*model.Message, error) {
	gacha := []string{
		"SSR",
		"SR",
		"R",
		"N",
	}
	result := gacha[randIntn(len(gacha))]
	return &model.Message{
		Username: "Gacha Bot",
		Body: result,
	}, nil
}

// Process はメッセージ本文からキーワードを抽出します
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
  params.Add("apikey", env.TalkAPIAppID)
  params.Add("query", text)

	post(talkAPIURLFormat, params, res)

	if res.Status != 0 {
		return nil, fmt.Errorf("%#v", res)
	}

	return &model.Message{
		Username: "Talk Bot",
		Body: res.Results[0].Reply,
	}, nil
}
