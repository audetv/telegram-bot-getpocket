package main

import (
	"github.com/audetv/telegram-bot-getpocket/internal/db/bbolt/tokenstore"
	"github.com/audetv/telegram-bot-getpocket/internal/pkg/repos/token"
	"github.com/audetv/telegram-bot-getpocket/internal/pkg/tgbot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"go.etcd.io/bbolt"
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

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := tokenstore.NewTokenRepository(db)

	telegramBot := tgbot.NewBot(bot, pocketClient, tokenRepository, "http://localhost/")
	err = telegramBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func initDB() (*bbolt.DB, error) {
	db, err := bbolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(token.AccessTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(token.RequestTokens))
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
