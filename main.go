package main

import (
	"os"

	"./logger"
	"./plugin"
	"./slack"
)

var log *logger.Logger

func init() {
	log = logger.GetLogger()
}

func main() {
	apiToken := os.Getenv("SLACK_API_TOKEN")
	if len(apiToken) == 0 {
		log.Error("SLACK_API_TOKEN is required.")
		return
	}

	session := slack.Session{Token: apiToken}
	err := session.Start()
	if err != nil {
		return
	}
	defer session.Close()

	registry := plugin.GetRegistry()
	registry.AddPlugin(&plugin.EchoPlugin{})
	registry.AddPlugin(&plugin.LoggerPlugin{})

	for {
		event, err := session.Receive()
		if err != nil {
			return
		}

		registry.Notify(&session, event)
	}
}
