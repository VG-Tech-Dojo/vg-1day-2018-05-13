package bot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/VG-Tech-Dojo/vg-1day-2018-05-13/NAKKA/env"
	"github.com/mrjones/oauth"
)

type Twitter struct {
	consumer    *oauth.Consumer
	accessToken *oauth.AccessToken
}

// get はurlにGETします
func get(url string, out interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return err
	}

	return nil
}

// post はurlにparamsをPOSTします
func post(url string, params url.Values, out interface{}) error {
	resp, err := http.PostForm(url, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, out)
	if err != nil {
		return err
	}

	return nil
}

// postJSON はinputをJSON形式でurlにPOSTします
func postJSON(url string, input interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return err
	}

	return nil
}

// randIntn は0からn-1までのintの乱数を返します
func randIntn(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

// TwitterAPI------------------------------------------------------

func NewTwitter(consumerKey, consumerSecret, accessToken, accessTokenSecret string) *Twitter {
	twitter := new(Twitter)
	twitter.consumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})
	twitter.accessToken = &oauth.AccessToken{accessToken, accessTokenSecret, make(map[string]string)}
	return twitter
}

func (t *Twitter) get(url string, params map[string]string) (interface{}, error) {
	response, err := t.consumer.Get(url, params, t.accessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// decode
	var result interface{}
	err = json.Unmarshal(b, &result)
	return result, err
}

func twitterGet(key string) (msg string, err error) {
	twitter := NewTwitter(env.TwitterConsumerKey, env.TwitterConsumerSecret, env.TwitterAccessToken, env.TwitterTokenSecret)

	// TwitterAPI検索結果を取得(落合陽一のツイート、リツイートなし、1件、最新のツイート)
	params := map[string]string{"q": "from:ochyai -rt", "count": "1", "result_type": "recent"}
	res, err := twitter.get(
		"https://api.twitter.com/1.1/search/tweets.json", // Resource URL
		params) // Parameters
	if err != nil {
		return "", err
	}

	msg = "[No tweet!!]"
	if statuses := res.(map[string]interface{})["statuses"]; statuses != nil {
		if status := statuses.([]interface{})[0]; status != nil {
			if text := status.(map[string]interface{})["text"]; text != nil {
				msg = text.(string)
			}
		}
	}

	return msg, nil
}
