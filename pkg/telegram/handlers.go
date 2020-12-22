package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	commandList  = "list"

	startMessage = "Привет! Чтобы начать сохранять ссылки в своем Pocket аккаунте, для начала тебе необходимо дать мне на это доступ. Для этого переходи по ссылке:\n%s"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandList:
		return b.handleListCommand()
	default:
		return b.handleUnknownCommand()
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	authLink, err := b.createAuthorizationLink(message.Chat.ID)
	if err != nil {
		return err
	}

	// TODO: store messages in config
	msgText := fmt.Sprintf(startMessage, authLink)
	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) handleListCommand() error {
	return nil
}

func (b *Bot) handleUnknownCommand() error {
	return nil
}

func (b *Bot) handleMessage() error {
	return nil
}

func (b *Bot) createAuthorizationLink(chatID int64) (string, error) {
	redirectUrl := b.generateRedirectURL(chatID)

	token, err := b.client.GetRequestToken(context.Background(), redirectUrl)
	if err != nil {
		return "", err
	}

	return b.client.GetAuthorizationURL(token, redirectUrl)
}

func (b *Bot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURL, chatID)
}
