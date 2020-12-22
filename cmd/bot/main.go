package main

import (
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/zhashkevych/go-pocket-sdk"
	"github.com/zhashkevych/telegram-pocket-bot/pkg/server"
	"github.com/zhashkevych/telegram-pocket-bot/pkg/storage"
	"github.com/zhashkevych/telegram-pocket-bot/pkg/storage/boltdb"
	"github.com/zhashkevych/telegram-pocket-bot/pkg/telegram"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	botApi, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	botApi.Debug = true

	pocketClient, err := pocket.NewClient(os.Getenv("CONSUMER_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	db, err := initBolt()
	if err != nil {
		log.Fatal(err)
	}
	storage := boltdb.NewTokenStorage(db)

	bot := telegram.NewBot(botApi, pocketClient, "http://localhost", storage)

	redirectServer := server.NewAuthServer("https://t.me/getpocket_client_bot", storage, pocketClient)

	go func() {
		if err := redirectServer.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initBolt() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Batch(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(storage.AccessTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(storage.RequestTokens))
		return err
	}); err != nil {
		return nil, err
	}

	return db, nil
}
