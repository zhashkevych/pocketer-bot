package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	invalidUrlError   = errors.New("url is invalid")
	unableToSaveError = errors.New("unable to save link to Pocket")
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	switch err {
	case invalidUrlError:
		messageText = b.messages.Errors.InvalidURL
	case unableToSaveError:
		messageText = b.messages.Errors.UnableToSave
	default:
		messageText = b.messages.Errors.Default
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.bot.Send(msg)
}
