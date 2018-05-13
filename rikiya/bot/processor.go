package bot

import (
	"regexp"
	"strings"

	"fmt"

	"net/url"

	"github.com/VG-Tech-Dojo/vg-1day-2018-05-13/rikiya/env"
	"github.com/VG-Tech-Dojo/vg-1day-2018-05-13/rikiya/model"
)

const (
	keywordAPIURLFormat      = "https://jlp.yahooapis.jp/KeyphraseService/V1/extract?appid=%s&sentence=%s&output=json"
	googlePlacesAPIURLFormat = "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=%f,%f&radius=%d&type=%s&key=%s&language=ja"
	GoogleMapsAPIURLFormat = "https://maps.googleapis.com/maps/api/geocode/json?key=%s&address=%s"
)

var places = [][]string{
	{
		"restaurant",
	},
	{
		"aquarium",
		"park",
		"museum",
		"shopping_mall",
	},
	{
		"cafe",
	},
}

type GooglePlacesAPIResult struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}


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

	GooglePlacesProcessor struct{}
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

func (p *GooglePlacesProcessor) Process(msgIn *model.Message) (*model.Message, error) {
	r := regexp.MustCompile("\\Adate (.*)\\z")
	matchedStrings := r.FindStringSubmatch(msgIn.Body)
	text := matchedStrings[1]

  // address -> lat,lng by Google maps api
	googleMapsAPIUrl := fmt.Sprintf(GoogleMapsAPIURLFormat, env.GoogleMapsAPIID, url.QueryEscape(text))
	googleMapsAPIRes := &struct {
		Results []struct {
			Geometry struct {
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
			} `json:"geometry"`
		} `json:"results"`
	}{}
	get(googleMapsAPIUrl, &googleMapsAPIRes)

  lat := googleMapsAPIRes.Results[0].Geometry.Location.Lat
  lng := googleMapsAPIRes.Results[0].Geometry.Location.Lng
	result_places := []string{}
	for _, v := range places {
		types := v[randIntn(len(v))]

		var googlePlacesAPIRes GooglePlacesAPIResult
		googlePlacesAPIUrl := fmt.Sprintf(googlePlacesAPIURLFormat, lat, lng, 2000, types, env.GooglePlacesAPIAppID)
		// fmt.Print(googlePlacesAPIUrl)
		get(googlePlacesAPIUrl, &googlePlacesAPIRes)
		// fmt.Print(googlePlacesAPIRes)
		spot := googlePlacesAPIRes.Results[randIntn(len(googlePlacesAPIRes.Results))]
		// fmt.Print(spot)
		result_places = append(result_places, spot.Name)
	}
	fmt.Print(result_places)
	msg := fmt.Sprintf(strings.Join(result_places, "->"))
	return &model.Message{
		Body: msg,
	}, nil
}
