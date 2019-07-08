package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("invalid arguments")
	}
	userName := os.Args[2]
	message := os.Args[3]

	api := slack.New(os.Args[1])
	users, _ := api.GetUsers()
	var userId string
	for _, user := range users {
		if (user.RealName == userName) || (user.Name == userName) {
			userId = user.ID
		}
	}
	_, _, channelID, _ := api.OpenIMChannel(userId)
	m := slack.MsgOptionText(message, true)
	api.SendMessage(channelID, m)
}
