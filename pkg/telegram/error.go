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
		messageText = "Это невалидная ссылка"
	case unableToSaveError:
		messageText = "Не удалось сохранить ссылку в Pocket. Попробуйте пожалуйста позже."
	default:
		messageText = "Произошла ошибка"
	}

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.bot.Send(msg)
}
