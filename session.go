package main

import (
	"fmt"

	websocket "golang.org/x/net/websocket"

	"./slack"
)

const WS_ORIGIN = "https://api.slack.com/"

type Session struct {
	Token  string
	wssUrl string
	conn   *websocket.Conn
}

func (session *Session) Start() error {
	var err error

	err = session.fetchWssUrl()
	if err != nil {
		return err
	}

	return session.startSession()
}

func (session *Session) Close() error {
	if session.conn == nil {
		return nil
	}

	err := session.conn.Close()
	if err != nil {
		fmt.Print(err)
	}

	return err
}

func (session *Session) Receive() (slack.Event, error) {
	event := slack.BaseEvent{}
	err := websocket.JSON.Receive(session.conn, &event)
	if err != nil {
		fmt.Print(err)
	}

	return event.ConcreteEvent(), err
}

func (session *Session) Send(event slack.Event) error {
	event.SetNextId()
	return websocket.JSON.Send(session.conn, event)
}

func (session *Session) fetchWssUrl() (err error) {
	req := slack.RtmStartApi{session.Token}
	session.wssUrl, err = req.WssUrl()
	if err != nil {
		fmt.Print("Session.fetchWssUrl:", err)
	}

	return err
}

func (session *Session) startSession() (err error) {
	session.conn, err = websocket.Dial(session.wssUrl, "", WS_ORIGIN)
	if err != nil {
		fmt.Print(err)
	}

	return err
}
