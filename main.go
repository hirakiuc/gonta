package main

import (
	"fmt"
	"os"

	"./plugin"
	"./slack"
)

func main() {
	apiToken := os.Getenv("SLACK_API_TOKEN")
	if len(apiToken) == 0 {
		fmt.Println("SLACK_API_TOKEN is required.")
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
