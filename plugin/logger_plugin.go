package plugin

import (
	"fmt"
	"time"

	"../slack"
)

type LoggerPlugin struct{}

func (plugin *LoggerPlugin) GetInstance() Plugin {
	return &LoggerPlugin{}
}

func (plugin *LoggerPlugin) IsAccept(event slack.Event) bool {
	return true
}

func (plugin *LoggerPlugin) Notify(session *slack.Session, event slack.Event) {
	fmt.Println(time.Now().Format(time.RFC3339), " : ", event)
}
