package plugin

import (
	"fmt"

	"../slack"
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
	fmt.Println("EchoPlugin:", event)
	evt, ok := event.(slack.MessageEvent)
	if ok == false {
		fmt.Println("EchoEvent:", ok, evt)
		return
	}

	msg := slack.BaseEvent{
		Id:      0,
		Type:    evt.Type,
		Channel: evt.Channel,
		Text:    "Hello !",
	}

	err := session.Send(msg)
	if err != nil {
		fmt.Println("EchoPlugin failed: ", err)
	} else {
		fmt.Println("EchoPlugin success.")
	}
}
