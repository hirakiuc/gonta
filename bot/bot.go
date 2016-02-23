package bot

import (
	"fmt"

	simpleJson "github.com/bitly/go-simplejson"
)

type SlackBot struct {
	Id   string
	Name string
}

var slackBot *SlackBot

func Initialize(json *simpleJson.Json) *SlackBot {
	if slackBot != nil {
		return slackBot
	}

	self := json.Get("self")

	slackBot = &SlackBot{
		Id:   self.Get("id").MustString(),
		Name: self.Get("name").MustString(),
	}
	return slackBot
}

func BotInstance() *SlackBot {
	return slackBot
}

func (bot *SlackBot) String() string {
	return fmt.Sprintf("SlackBot<Id: %s, Name: %s>", bot.Id, bot.Name)
}
