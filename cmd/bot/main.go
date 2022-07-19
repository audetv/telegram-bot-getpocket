package main

import (
	"github.com/audetv/telegram-bot-getpocket/pkg/tgbot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(":token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := tgbot.NewBot(bot)
	err = telegramBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
