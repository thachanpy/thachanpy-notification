package main

import (
	"log"
	"github.com/thachanpy/thachanpy-notification/slack"
)

func main() {
	err := slack.SendSlackMessage()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Sent to Slack successfully")
	}
}