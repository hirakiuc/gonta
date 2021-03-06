package plugin

import (
	"fmt"

	"github.com/hirakiuc/gonta/internal/slack"
)

// EchoPlugin provides Echo response.
type EchoPlugin struct{}

// GetInstance return a new EchoPlugin instance.
func (plugin *EchoPlugin) GetInstance() Plugin {
	return &EchoPlugin{}
}

// IsAccept check whether this plugin accept the slack.Event or not.
func (plugin *EchoPlugin) IsAccept(event slack.Event) bool {
	if event.EventType() == "message" {
		return true
	}

	return false
}

// Notify the event to this EchoPlugin.
func (plugin *EchoPlugin) Notify(session *slack.Session, event slack.Event) {
	log.Debug("EchoPlugin: %v", event)
	evt, ok := event.(slack.MessageEvent)
	if ok == false {
		log.Debug("EchoEvent: %v %v", ok, evt)
		return
	}

	response := slack.BaseEvent{
		Id:      0,
		Type:    evt.Type,
		Channel: evt.Channel,
		Text:    fmt.Sprintf("Hello ! I'm %s.", session.Bot.Name),
	}

	err := session.Send(response)
	if err != nil {
		log.Error("EchoPlugin failed: %v", err)
	}

	log.Info("EchoPlugin success.")
}
