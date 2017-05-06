package slack

import (
	"fmt"
	"sync/atomic"
)

var counter uint64

type BaseEvent struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type Event interface {
	String() string
	ConcreteEvent() Event
	SetNextId()

	EventType() string
}

func (event BaseEvent) String() string {
	return fmt.Sprintf(
		"{ Id: %d, Type: %s, Channel: %s, Text: %s }",
		event.Id, event.Type, event.Channel, event.Text)
}

func (event BaseEvent) ConcreteEvent() Event {
	switch event.Type {
	case "message":
		return MessageEvent{&event}
	default:
		return event
	}
}

func (event BaseEvent) SetNextId() {
	event.Id = atomic.AddUint64(&counter, 1)
}

func (event BaseEvent) EventType() string {
	return event.Type
}

type MessageEvent struct {
	*BaseEvent
}
