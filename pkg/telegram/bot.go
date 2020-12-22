package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
	"github.com/zhashkevych/telegram-pocket-bot/pkg/storage"
	"log"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	client      *pocket.Client
	redirectURL string

	storage   storage.TokenStorage
}

func NewBot(bot *tgbotapi.BotAPI, client *pocket.Client, redirectURL string, storage storage.TokenStorage) *Bot {
	return &Bot{
		bot:         bot,
		client:      client,
		redirectURL: redirectURL,
		storage:     storage,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Handle commands
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				log.Println("handle command error", err.Error())
			}

			continue
		}

		// Handle regular messages
		if err := b.handleMessage(update.Message); err != nil {

		}
	}

	return nil
}