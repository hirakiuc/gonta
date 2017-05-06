package slack

import (
	websocket "golang.org/x/net/websocket"

	"github.com/hirakiuc/gonta/internal/bot"
	"github.com/hirakiuc/gonta/internal/logger"
)

const WS_ORIGIN = "https://api.slack.com/"

var log *logger.Logger

type Session struct {
	Token  string
	Bot    *bot.SlackBot
	wssUrl string
	conn   *websocket.Conn
}

func init() {
	log = logger.GetLogger()
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
		log.Error("Failed to close Websocket connection: %v", err)
	}

	return err
}

func (session *Session) Receive() (Event, error) {
	event := BaseEvent{}
	err := websocket.JSON.Receive(session.conn, &event)
	if err != nil {
		log.Error("Failed to read next event: %v", err)
	}

	return event.ConcreteEvent(), err
}

func (session *Session) Send(event Event) error {
	event.SetNextId()
	return websocket.JSON.Send(session.conn, event)
}

func (session *Session) fetchWssUrl() (err error) {
	req := RtmStartApi{session.Token}
	session.wssUrl, err = req.WssUrl()
	if err != nil {
		log.Error("Failed to fetch WssUrl: %v", err)
	}

	return err
}

func (session *Session) startSession() (err error) {
	session.conn, err = websocket.Dial(session.wssUrl, "", WS_ORIGIN)
	if err != nil {
		log.Error("Failed to connect Websocket: %v", err)
	}

	session.Bot = bot.BotInstance()

	return err
}
