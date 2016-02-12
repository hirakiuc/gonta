package slack

import (
	"fmt"
	"io/ioutil"
	"net/http"

	simpleJson "github.com/bitly/go-simplejson"
)

type RtmStartApi struct {
	Token string
}

const BASE_API = "https://slack.com/api/rtm.start"

func (api *RtmStartApi) requestUrl() string {
	return fmt.Sprintf(BASE_API+"?token=%s", api.Token)
}

func (api *RtmStartApi) WssUrl() (url string, err error) {
	res, err := http.Get(api.requestUrl())
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	js, err := simpleJson.NewJson(body)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	wss, err := js.Get("url").String()
	if err != nil {
		fmt.Print(err)
		return
	}

	return wss, nil
}
