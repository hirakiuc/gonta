package slack

import "fmt"

type Event struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func (event Event) String() string {
	return fmt.Sprintf(
		"{ Id: %d, Type: %s, Channel: %s, Text: %s }",
		event.Id, event.Type, event.Channel, event.Text)
}
