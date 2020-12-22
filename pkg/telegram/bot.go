package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	client      *pocket.Client
	redirectURL string
}

func NewBot(bot *tgbotapi.BotAPI, client *pocket.Client, redirectURL string) *Bot {
	return &Bot{bot: bot, client: client, redirectURL: redirectURL}
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
			b.handleCommand(update.Message)
			continue
		}

		// Handle regular messages
		b.handleMessage()
	}

	return nil
}
