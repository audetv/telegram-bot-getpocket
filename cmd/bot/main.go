package main

import (
	"github.com/audetv/telegram-bot-getpocket/pkg/tgbot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(":token")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient(":consumerKey")
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := tgbot.NewBot(bot, pocketClient, "http://localhost/")
	err = telegramBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
