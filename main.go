package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nlopes/slack"
)

func SmsToSlackHandler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Printf("req.ParseForm failed: %v", err)
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("req.ParseForm failed: %v", err)))
		return
	}

	slackToken := os.Getenv("SLACK_API_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")

	slackClient := slack.New(slackToken)

	message := fmt.Sprintf("TwilioからのSMS: \n差出人: %s\n本文: %s", req.FormValue("From"), req.FormValue("Body"))
	_, _, err := slackClient.PostMessage(
		channelID,
		slack.MsgOptionText(string(message), false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Printf("failed to post message to Slack: %v", err)
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("failed to post message to Slack: %v", err)))
		return
	}

	w.WriteHeader(200)
}
