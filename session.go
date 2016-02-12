package main

import (
	"fmt"
	"sync/atomic"

	websocket "golang.org/x/net/websocket"

	"./slackapi"
)

const WS_ORIGIN = "https://api.slack.com/"

type Session struct {
	Token  string
	wssUrl string
	conn   *websocket.Conn
}

type Message struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
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

func (session *Session) Receive(msg *Message) error {
	return websocket.JSON.Receive(session.conn, msg)
}

var counter uint64

func (session *Session) Send(msg Message) error {
	msg.Id = atomic.AddUint64(&counter, 1)
	return websocket.JSON.Send(session.conn, msg)
}

func (session *Session) fetchWssUrl() (err error) {
	req := slackapi.RtmStartApi{session.Token}
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
