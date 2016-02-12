package main

import (
	"fmt"
	"os"
)

func main() {
	apiToken := os.Getenv("SLACK_API_TOKEN")
	if len(apiToken) == 0 {
		fmt.Println("SLACK_API_TOKEN is required.")
		return
	}

	session := Session{Token: apiToken}

	err := session.Start()
	if err != nil {
		return
	}
	defer session.Close()

	event := Event{}
	for {
		if err := session.Receive(&event); err != nil {
			return
		}

		fmt.Println(event)
	}
}
