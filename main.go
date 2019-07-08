package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

func showHelp() {
	fmt.Printf("Use `report [token] [username] [message]` to send the message\n")
}

func isNoOptionProvided() bool {
	return len(os.Args) < 2
}

func isOptionHelp() bool {
	return len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help")
}

func main() {
	if isNoOptionProvided() || isOptionHelp() {
		showHelp()
		return
	}

	if len(os.Args) < 4 {
		log.Fatal("invalid number of arguments")
	}
	userName := os.Args[2]
	message := os.Args[3]

	api := slack.New(os.Args[1])
	users, err := api.GetUsers()
	if err != nil {
		log.Fatalf("Not able to get users due to ==> %s", err)
	}

	var userId string
	for _, user := range users {
		if (user.RealName == userName) || (user.Name == userName) {
			userId = user.ID
		}
	}
	_, _, channelID, err := api.OpenIMChannel(userId)
	if err != nil {
		log.Fatalf("Not able to open direct channel due to ==> %s", err)
	}

	m := slack.MsgOptionText(message, true)
	_, _, _, err = api.SendMessage(channelID, m)
	if err != nil {
		log.Fatalf("Not able to send message due to ==> %s", err)
	}
}
